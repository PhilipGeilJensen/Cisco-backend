package cisco

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
)

type Connection struct {
	*ssh.Client
}


func Connect(addr, user, password string) (*Connection, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{conn}, nil

}

func (conn *Connection) SendCommands(cmd []string) ([]byte, error) {
	// Create sesssion
	sess, err := conn.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer sess.Close()

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// send the commands
	for _, cmd := range cmd {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
	return []byte("ok"), err
}


func (conn *Connection) SendShowCommand(cmd string) ([]byte, error) {
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return []byte{}, err
	}

	output, err := session.Output(cmd)
	if err != nil {
		return output, fmt.Errorf("failed to execute command '%s' on server: %v", cmd, err)
	}

	return output, err
}