package ssh

import (
	"bytes"
	"errors"

	"github.com/sirupsen/logrus"
)

type SingleCmd struct {
	parameters Parameters
}

const (
	srcUserField = "src_user"
)

func NewSingleCmd(sshParameters Parameters) *SingleCmd {
	ret := &SingleCmd{
		parameters: sshParameters,
	}
	return ret
}

func (selfPtr *SingleCmd) Execute(cmd string) (output string, err error) {
	if err := selfPtr.parameters.verify(); err != nil {
		logrus.Error(err)
		return "error output", err
	}
	session, err := selfPtr.parameters.createSSHSession()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer session.Close()
	var stdOutBuf, stdErrBuf bytes.Buffer
	session.Stdout = &stdOutBuf
	session.Stderr = &stdErrBuf
	session.Run(cmd)
	if stdErrBuf.String() != "" {
		return stdOutBuf.String(), errors.New(stdErrBuf.String())
	}
	return stdOutBuf.String(), nil
}
