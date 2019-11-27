package options

import (
	"flag"
	"fmt"
	"log"
)

type Opts struct {
	RedisURL        string
	Port            int64
	StorageKeyName  string
	EmptyIdentifier string
}

var GlobalConfig *Opts

func ParseFlags() {
	var redisURL string
	var port int64
	var storageKeyName string
	var emptyIdentifier string

	flag.StringVar(&redisURL, "redis-url", "", "Redis DB URL")
	flag.Int64Var(&port, "port", 44225, "Port")
	flag.StringVar(&storageKeyName, "storage-key-name", "", "Storage Key name for saving passport headers")
	flag.StringVar(&emptyIdentifier, "empty-identifier", "unknown", "Identifier for empty values")
	flag.Parse()

	GlobalConfig = &Opts{
		RedisURL:        redisURL,
		Port:            port,
		StorageKeyName:  storageKeyName,
		EmptyIdentifier: emptyIdentifier,
	}

	if GlobalConfig.RedisURL == "" {
		log.Panicf("Redis URL not specified")
	}

	fmt.Printf("Application options %+v", GlobalConfig)
}
