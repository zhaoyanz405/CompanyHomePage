package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func handleWriteBack(c net.Conn) {
	defer c.Close()
	io.Copy(c, c)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		log.Print("client has accept, addr: ", conn.RemoteAddr())

		if err != nil {
			log.Print(err)
			continue
		}
		//go handleConn(conn)
		go handleWriteBack(conn)
		log.Print("client has close, addr: ", conn.RemoteAddr())
	}
}
