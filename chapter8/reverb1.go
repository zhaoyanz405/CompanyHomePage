package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func repeat(c net.Conn) {
	defer func() {
		log.Print("client has close, addr: ", c.RemoteAddr())
		c.Close()
	}()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
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
		go repeat(conn)
	}
}
