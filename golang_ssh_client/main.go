package main

import (
	"fmt"
	"myssh/ssh"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
)

const (
	user    = "vagrant"
	host    = "localhost"
	port    = 2222
	keyPath = "/home/frankie/git_repo/code_4_blog/golang_ssh_client/test/.vagrant/machines/default/virtualbox/private_key"
)

func main() {
	f, err := os.OpenFile("ssh.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(f)
	logrus.SetLevel(logrus.DebugLevel)
	parameters := ssh.Parameters{
		TargetUser:     user,
		RemoteHost:     host,
		RemoteHostPort: port,
		PrivateKeyPath: keyPath,
	}
	// single ssh cmd
	sshCmd := ssh.NewSingleCmd(parameters)
	sshCmd.Execute("touch /tmp/test_ssh")
	session := ssh.NewInteractiveSession(
		ssh.InteractiveSessionConfig{
			SessionID:  uuid.NewV4().String(),
			Parameters: parameters,
			ExpiryTime: time.Now().Add(10 * time.Second),
		},
	)
	session.DoSSH()
}
