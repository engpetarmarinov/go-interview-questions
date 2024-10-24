package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	wgServerReady := &sync.WaitGroup{}
	wgServerReady.Add(1)
	done := make(chan struct{})
	go func() {
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			panic(err)
		}

		defer l.Close()

		wgServerReady.Done()

		for {
			c, err := l.Accept()
			if err != nil {
				log.Println(err)
			}

			content, err := io.ReadAll(c)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(string(content))

			c.Write([]byte("pong"))

			//c.Close()
			done <- struct{}{}
		}
	}()

	wgServerReady.Wait()

	connPool := preWarmConnectionsToService()

	for i := 0; i < 100; i++ {
		serviceConn := connPool.Get()
		sc, ok := serviceConn.(net.Conn)
		if !ok {
			log.Println("service conn is not net.Conn")
			return
		}

		_, err := sc.Write([]byte("ping"))
		if err != nil {
			log.Println(err)
			return
		}
		connPool.Put(serviceConn)
	}

	<-done

}

func preWarmConnectionsToService() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}

	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func connectToService() any {
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
		return nil
	}

	return c
}
