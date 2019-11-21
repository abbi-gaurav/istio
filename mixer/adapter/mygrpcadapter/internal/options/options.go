package options

import (
	"flag"
	"fmt"
	"log"
)

type Opts struct {
	RedisURL        string
	Port            int64
	PassportHeaders string
	StorageKeyName  string
	DimensionsKey   string
	EmptyIdentifier string
}

var GlobalConfig *Opts

func ParseFlags() {
	var redisURL string
	var port int64
	var passportHeaders string
	var storageKeyName string
	var dimensionsKey string
	var emptyIdentifier string

	flag.StringVar(&redisURL, "redis-url", "", "Redis DB URL")
	flag.Int64Var(&port, "port", 44225, "Port")
	flag.StringVar(&passportHeaders, "passport-headers", "", "Comma separated list of passport headers")
	flag.StringVar(&storageKeyName, "storage-key-name", "", "Storage Key name for saving passport headers")
	flag.StringVar(&dimensionsKey, "dimensions-key", "header_name", "Dimensions Key to get header names")
	flag.StringVar(&emptyIdentifier, "empty-identifier", "unknown", "Identifier for empty values")
	flag.Parse()

	GlobalConfig = &Opts{
		RedisURL:        redisURL,
		Port:            port,
		PassportHeaders: passportHeaders,
		StorageKeyName:  storageKeyName,
		DimensionsKey:   dimensionsKey,
		EmptyIdentifier: emptyIdentifier,
	}

	if GlobalConfig.RedisURL == "" {
		log.Panicf("Redis URL not specified")
	}

	fmt.Printf("Application options %+v", GlobalConfig)
}
