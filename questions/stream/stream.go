package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	socketPath := "./unix.sock:w" +
		""
	if err := os.RemoveAll(socketPath); err != nil {
		log.Fatalf("remove error: %v", err)
	}

	err = syscall.Bind(fd, &syscall.SockaddrUnix{
		Name: socketPath,
	})
	if err != nil {
		panic(err)
	}

	err = syscall.Listen(fd, syscall.SOMAXCONN)
	if err != nil {
		panic(err)
	}

	for {
		cfd, sockAddr, err := syscall.Accept(fd)
		if err != nil {
			panic(err)
		}
		if addr, ok := sockAddr.(*syscall.SockaddrUnix); ok {
			fmt.Printf("Client connected from %s\n", addr.Name)
		} else {
			fmt.Println("Client connected from unknown address")
		}

		go handleConnection(cfd)
	}
}

func handleConnection(cfd int) {
	buf := make([]byte, 1024)

	n, err := syscall.Read(cfd, buf)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	msg := buf[:n]
	syscall.Write(cfd, msg)
	defer syscall.Close(cfd)
}
