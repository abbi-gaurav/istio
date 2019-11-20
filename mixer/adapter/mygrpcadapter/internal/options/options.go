package options

import (
	"flag"
	"fmt"
	"log"
)

type Opts struct {
	RedisURL string
	Port     int64
}

var GlobalConfig *Opts

func ParseFlags() {
	var redisURL string
	var port int64

	flag.StringVar(&redisURL, "redis-url", "", "Redis DB URL")
	flag.Int64Var(&port, "port", 44225, "Port")
	flag.Parse()

	GlobalConfig = &Opts{
		RedisURL: redisURL,
		Port:     port,
	}

	if GlobalConfig.RedisURL == "" {
		log.Panicf("Redis URL not specified")
	}

	fmt.Printf("Application options %+v", GlobalConfig)
}
