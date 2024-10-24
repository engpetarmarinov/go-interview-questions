package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	sockPath := "unix.sock"

	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatalf("socket error: %v", err)
	}

	if err := os.RemoveAll(sockPath); err != nil {
		log.Fatalf("remove error: %v", err)
	}

	sa := &syscall.SockaddrUnix{Name: sockPath}

	if err := syscall.Bind(fd, sa); err != nil {
		log.Fatalf("bind error: %v", err)
	}

	fmt.Printf("Listening for datagrams on %s...\n", sockPath)

	buf := make([]byte, 1024)

	for {
		n, from, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.Fatalf("recvfrom error: %v", err)
		}

		fmt.Printf("Received %d bytes from %s: %s\n", n, from.(*syscall.SockaddrUnix).Name, string(buf[:n]))

		response := "Acknowledged: " + string(buf[:n])
		syscall.Sendto(fd, []byte(response), 0, from)
	}
}
