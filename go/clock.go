// Adaptado de Alan A. A. Donovan & Brian W. Kernighan.
// a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	//escuta na porta 8000 (pode ser monitorado com lsof -Pn -i4 | grep 8000)
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//aceita uma conexão criada por um cliente
		conn, err := listener.Accept()
		if err != nil {
			// falhas na conexão. p.ex abortamento
			log.Print(err)
			continue
		}
		// serve a conexão estabelecida
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	defer c.Close()
	for {
		// envia o conteúdo servido na conexão
		_, err := io.WriteString(c, time.Now().Format("02:05:00\n"))

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

//!-
