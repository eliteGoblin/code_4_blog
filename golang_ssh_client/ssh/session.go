package ssh

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	officialSSH "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type InteractiveSession struct {
	config        InteractiveSessionConfig
	timeExpiryChn chan int
	exitChn       chan int
}

type InteractiveSessionConfig struct {
	SessionID  string
	Parameters Parameters
	ExpiryTime time.Time
}

func (selfPtr *InteractiveSessionConfig) verify() error {
	if err := selfPtr.Parameters.verify(); err != nil {
		return err
	}
	if selfPtr.SessionID == "" {
		logrus.Warn("empty session id")
	}
	return nil
}

const (
	placeHolder             = "OUTPUT"
	checkTickerIntervalSecs = 2
)

func NewInteractiveSession(config InteractiveSessionConfig) *InteractiveSession {
	session := &InteractiveSession{
		config: config,
	}
	session.timeExpiryChn = make(chan int, 1)
	session.exitChn = make(chan int, 1)
	return session
}

// DoSSH will block until user end the SSH session
func (selfPtr *InteractiveSession) DoSSH() {
	logrus.Debugf("interactiveSession config %+v", selfPtr.config)
	if err := selfPtr.config.verify(); err != nil {
		logrus.Error(err)
		return
	}
	if err := CheckExpired(selfPtr.config.ExpiryTime); err != nil {
		logrus.Error(err)
		return
	}
	session, err := selfPtr.config.Parameters.createSSHSession()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer session.Close()
	// config and run session
	err = selfPtr.configSessionAndRun(session)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func (selfPtr *InteractiveSession) startExpiryChecker(duration time.Duration) {
	checker := time.NewTicker(duration)
	go func() {
		for {
			<-checker.C
			if err := CheckExpired(selfPtr.config.ExpiryTime); err != nil {
				logrus.Infof("session finish: %s signal exit...", err)
				selfPtr.timeExpiryChn <- 1 // notify session finished
				return
			}
		}
	}()
}

func CheckExpired(expiryTime time.Time) error {
	if time.Now().After(expiryTime) {
		return fmt.Errorf("expired: %s; now: %s",
			expiryTime.Format(time.RFC3339),
			time.Now().Format(time.RFC3339))
	}
	return nil
}

func (selfPtr *InteractiveSession) configSessionAndRun(session *officialSSH.Session) error {
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr
	outPipe, errOut := session.StdoutPipe()
	if errOut != nil {
		return fmt.Errorf("%s", errOut)
	}
	modes := officialSSH.TerminalModes{
		officialSSH.ECHO:          1,
		officialSSH.ECHOCTL:       0,
		officialSSH.TTY_OP_ISPEED: 14400,
		officialSSH.TTY_OP_OSPEED: 14400,
	}
	termFD := int(os.Stdin.Fd())
	width, height, err := terminal.GetSize(termFD)
	if err != nil {
		return err
	}
	termState, err := terminal.MakeRaw(termFD)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			logrus.Warnf("recover: %+v", r)
		}
		terminal.Restore(termFD, termState)
	}()
	// start routine to check if expiry
	selfPtr.startExpiryChecker(checkTickerIntervalSecs * time.Second)
	err = session.RequestPty("xterm-256color", width, height, modes)
	if err != nil {
		return err
	}
	err = session.Shell()
	if err != nil {
		return err
	}
	// set resize
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			logrus.Info("SIGWINCH signal received")
			width, height, err := terminal.GetSize(termFD)
			if err != nil {
				logrus.Warn(err.Error())
			} else {
				if err := session.WindowChange(height, width); err != nil {
					logrus.Warn(err.Error())
				}
			}
		}
	}()
	ch <- syscall.SIGWINCH // Initial resize.
	peepHoleR, peepHoleW := io.Pipe()
	multiWriterOut := io.MultiWriter(os.Stdout, peepHoleW)
	outReader := bufio.NewReader(peepHoleR)
	go func() {
		line, errReadLog := outReader.ReadString('\r')
		for errReadLog == nil {
			logrus.Infof("%s|%s", placeHolder, string(line))
			line, errReadLog = outReader.ReadString('\r')
		}
	}()
	go func() {
		_, err = io.Copy(multiWriterOut, outPipe)
	}()
	go func() {
		session.Wait()
		selfPtr.exitChn <- 1
	}()
	select {
	case <-selfPtr.timeExpiryChn:
		logrus.Info("time expiry msg received, exit...")
		return nil
	case <-selfPtr.exitChn:
		return nil
	}
}
