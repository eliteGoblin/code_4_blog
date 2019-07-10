package ssh

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	officialSSH "golang.org/x/crypto/ssh"
)

type Parameters struct {
	TargetUser     string
	RemoteHost     string
	RemoteHostPort int
	PrivateKeyPath string
}

func (selfPtr *Parameters) verify() error {
	if selfPtr.TargetUser == "" ||
		selfPtr.RemoteHost == "" ||
		selfPtr.RemoteHostPort == 0 ||
		selfPtr.PrivateKeyPath == "" {
		return fmt.Errorf("invalid config %+v", *selfPtr)
	}
	return nil
}

func (selfPtr *Parameters) createSSHSession() (session *officialSSH.Session, err error) {
	clientConfig, err := selfPtr.generateSSHClientConfig()
	session = &officialSSH.Session{}
	if err != nil {
		return session, err
	}
	addr := formatSSHDialStr(selfPtr.RemoteHost, selfPtr.RemoteHostPort)
	client, err := officialSSH.Dial("tcp", addr, clientConfig)
	if err != nil {
		return session, err
	}
	// create new session
	return client.NewSession()
}

func (selfPtr *Parameters) generateSSHClientConfig() (session *officialSSH.ClientConfig, err error) {
	auth := make([]officialSSH.AuthMethod, 0)
	pemBytes, err := ioutil.ReadFile(selfPtr.PrivateKeyPath)
	if err != nil {
		return &officialSSH.ClientConfig{}, err
	}
	var signer officialSSH.Signer
	signer, err = officialSSH.ParsePrivateKey(pemBytes)
	if err != nil {
		return &officialSSH.ClientConfig{}, err
	}
	auth = append(auth, officialSSH.PublicKeys(signer))
	if err != nil {
		return &officialSSH.ClientConfig{}, err
	}
	auth = append(auth, officialSSH.PublicKeys(signer))
	config := officialSSH.Config{
		Ciphers: []string{
			"aes128-ctr",
			"aes192-ctr",
			"aes256-ctr",
			"aes128-gcm@openssh.com",
			"arcfour256",
			"arcfour128",
			"aes128-cbc",
			"3des-cbc",
			"aes192-cbc",
			"aes256-cbc",
		},
	}
	return &officialSSH.ClientConfig{
		User:    selfPtr.TargetUser,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key officialSSH.PublicKey) error {
			return nil
		},
	}, nil
}

func formatSSHDialStr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
