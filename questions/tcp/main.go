package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type Session struct {
	Authenticated bool
}

type cmd string

const (
	cmdEXIT cmd = "EXIT"
	cmdPING cmd = "PING"
	cmdBEAT cmd = "BEAT"
	cmdAUTH cmd = "AUTH"
)

func main() {
	l, err := net.Listen("unix", "./unix.sock")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(c.RemoteAddr(), "connected")
	defer c.Close()
	err := c.SetDeadline(time.Now().Add(60 * time.Second))
	if err != nil {
		log.Println(err)
		return
	}

	ctx := context.WithValue(context.Background(), "session", &Session{Authenticated: false})
	for {
		command, err := readCmd(c)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("Command:", command)
		data, err := readData(c)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("Data:", data)
		go handleCmd(ctx, c, command, data)
	}
}

func handleCmd(ctx context.Context, conn net.Conn, command string, data string) {
	fmt.Printf("%s: Session: %v\n", conn.RemoteAddr(), ctx.Value("session").(*Session))
	command = strings.ToUpper(command)
	switch cmd(command) {
	case cmdEXIT:
		conn.Write([]byte("Exiting\n"))
		conn.Close()
		return
	case cmdPING:
		conn.Write([]byte("PONG\n"))
		conn.Write([]byte(data))
		return
	case cmdBEAT:
		conn.Write([]byte("Heart beat received. Extending connection deadline.\n"))
		conn.SetDeadline(time.Now().Add(60 * time.Second))
	case cmdAUTH:
		err := authToken(data)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			return
		}

		ctx.Value("session").(*Session).Authenticated = true
		conn.Write([]byte("Authenticated successfully\n"))
		return
	default:
		log.Println("Unknown command:", command)
	}
}

func authToken(data string) error {
	data = strings.TrimSpace(data)
	fmt.Println(data, len(data), len("securitytoken"))
	if data == "securitytoken" {
		return nil
	}

	return errors.New("invalid token")
}

func readCmd(c net.Conn) (string, error) {
	command := make([]byte, 4)
	_, err := c.Read(command)
	if err != nil {
		return "", err
	}

	receivedCmd := string(command)
	return receivedCmd, err
}

func readData(conn net.Conn) (string, error) {
	var data string
	dataBuf := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(dataBuf)
		if err != nil {
			if err != io.EOF {
				break
			}

			return "", err
		}

		data = data + string(dataBuf[:bytesRead])
		if bytesRead < len(dataBuf) {
			break
		}
	}

	return data, nil
}
