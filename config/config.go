package config

import (
	"flag"
	"sync"
)

type Config struct {
	Port    string
	BaseURL string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		port := flag.String("port", "8080", "Port to listen on for the web server")
		isHttps := flag.String("https", "false", "Protocol of base url")
		baseUrl := flag.String("redirect", "kuryltai.kz:8000", "Domen of base url")
		flag.Parse()
	
		if *isHttps == "true" {
			*baseUrl = "https://" + *baseUrl
		} else {
			*baseUrl = "http://" + *baseUrl
		}
	
		cfg = Config{
			Port:    *port,
			BaseURL: *baseUrl,
		}
	})

	return &cfg
}
