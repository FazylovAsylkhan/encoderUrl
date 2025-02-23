package main

import (
	"github.com/FazylovAsylkhan/encoderUrl/cmd/server"
	"github.com/FazylovAsylkhan/encoderUrl/internal/config"
)

func main() {
	cfg := config.Get()
	server.Start(cfg)	
}
