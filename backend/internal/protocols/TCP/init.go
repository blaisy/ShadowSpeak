// Package tcp represents a TCP connection recived from a remote client connection
// and handle it appropriately
//
// @author Gabriel Moraga Figueroa, Elia Todescato
// @version 0.1-beta
//
// Copyright (c) 2024 MIT Licensed
package tcp

import (
	"net"
	"os"

	"github.com/charmbracelet/log"
)

// Start a new TCP connection
func Start() { listenIncomingConn(getClient()) }

// getClient returns a TCP client
func getClient() net.Listener {

	addr := "localhost"
	port := os.Getenv("PORT")

	address := addr + ":" + port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("Error creating tcp listener: ", err)
		os.Exit(1)
	}

	log.Info("Listening on uri ", "addr", address)
	return listener
}

// listenIncomingConn connects the given TCP connection to the server
func listenIncomingConn(c net.Listener) {
	for {
		conn, err := c.Accept()
		if err != nil {
			log.Error("Error accepting connection: ", err)
			continue
		}
		go tcpHandler(conn)
	}
}

// tcpHandler handle connection
func tcpHandler(conn net.Conn) {
	defer conn.Close()
	log.Info("TCP connection established on ", "addr", conn.RemoteAddr().String())
}
