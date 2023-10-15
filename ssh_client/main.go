package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	flag.String("target", "127.0.0.1", "Target to run a command")
}

func run(target, cmd string) (string, error) {
	cfg := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}
	_ = cfg
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:22"), cfg)
	defer client.Close()
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
