package main

import (
	"github.com/FazylovAsylkhan/encoderUrl/cmd/httpService"
	"github.com/FazylovAsylkhan/encoderUrl/config"
)

func main() {
	cfg := config.Get()
	httpService.Start(cfg)	
}
