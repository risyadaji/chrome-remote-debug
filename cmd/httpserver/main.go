package main

import (
	_ "github.com/payfazz/chrome-remote-debug/config/env"

	"github.com/payfazz/chrome-remote-debug/internal/httpserver"
)

func main() {
	// start api server
	s := httpserver.NewServer()
	s.Serve("9000")
}
