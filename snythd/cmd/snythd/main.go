package main

import (
	"flag"
	"fmt"

	"github.com/transcriptaze/snyth/snythd/httpd"
)

var VERSION = "v0.1.0"

var options = struct {
	html    string
	address string
	port    uint
	debug   bool
}{
	html:    "./html",
	address: "0.0.0.0",
	port:    9000,
	debug:   false,
}

func main() {
	fmt.Printf("snythd - %v\n", VERSION)

	flag.StringVar(&options.address, "address", options.address, "HTTP server address (defaults to 0.0.0.0)")
	flag.UintVar(&options.port, "port", options.port, "HTTP port (defaults to 9000")
	flag.StringVar(&options.html, "html", options.html, "HTML folder (requires debug )")
	flag.BoolVar(&options.debug, "debug", options.debug, "enables internal debug mode")
	flag.Parse()

	httpd.Run(options.html, options.address, options.port, options.debug)
}
