// Package main is the entry point for the ShadowSpeak application.
//
// @author Gabriel Moraga Figueroa, Elia Todescato
// @version 0.1-beta
//
// @example cmd/main --mode=dev --port=8080
//
// Copyright (c) 2024 MIT Licensed
package main

import (
	"flag"
	"fmt"
	"os"

	tcp "github.com/blaisy/ShadowSpeak.git/internal/protocols/TCP"
	"github.com/charmbracelet/log"
)

// ShadowSpeak variables used to configure the Global variables
var (
	mode string
	port int
)

// main is the entry point for the ShadowSpeak application
func main() {
	setParams()

	tcp.Start()
}

// setParams sets the global mode variable and port number
func setParams() {
	// set production mode
	flag.StringVar(
		&mode,
		"mode",
		"prod",
		"mode rappresent the mode of the current process it can be used in two ways: dev or prod mode (default mode)",
	)

	// set port number (default port number is 8080)
	flag.IntVar(
		&port,
		"port",
		8080,
		"port rappresent the port where the process is running (default port is 8080)",
	)

	if mode == "" || port <= 0 {
		log.Error("Can't set production mode or port number MODE: %s, PORT: %d", mode, port)
		os.Exit(1)
	}

	os.Setenv("MODE", mode)
	os.Setenv("PORT", fmt.Sprint(port))
}
