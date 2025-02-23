package config

import (
	"flag"
	"sync"
)

type Config struct {
	Address    string
	BaseURL string
}

var (
	cfg  Config
	once sync.Once
)


func Get() *Config {
	once.Do(func() {
		address := flag.String("address", "localhost:8080", "Address to listen on for the web server")
		isHttps := flag.String("https", "false", "Protocol of base url")
		baseUrl := flag.String("baseURL", "kuryltai.kz:8000", "Base url for redirect")
		flag.Parse()
	
		if *isHttps == "true" {
			*baseUrl = "https://" + *baseUrl
		} else {
			*baseUrl = "http://" + *baseUrl
		}
	
		cfg = Config{
			Address:    *address,
			BaseURL: *baseUrl,
		}
	})

	return &cfg
}
