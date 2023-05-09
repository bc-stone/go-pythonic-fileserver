/*
A Pythonic-style HTTP file server.
By default, the server listens on all IP addresses present on the
system ('0.0.0.0') and is bound to port 8000.  The current directory
('.') is shared, relative to the location of 'fileserver.go' or the
binary if it is compiled.

Usage:

go run fileserver.go [optional flags]

The flags are:

-p
The TCP port to listen on.  Defaults to 8000.

-i
The interface IP address to use.  Defaults to 0.0.0.0.

-d
The directory path to share.  Defaults to '.'
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8000", "the port of http file server")
	var addr string
	flag.StringVar(&addr, "i", "0.0.0.0", "the IP address of the file server")
	var dir string
	flag.StringVar(&dir, "d", ".", "the directory path to share")
	flag.Parse()

	fmt.Printf("Serving HTTP on %s port %s at location '%s' ...\n", addr, port, dir)
	h := http.FileServer(http.Dir(dir))
	http.ListenAndServe(addr+":"+port, h)
}
