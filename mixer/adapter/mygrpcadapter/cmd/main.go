package main

import (
	"fmt"
	"istio.io/istio/mixer/adapter/mygrpcadapter"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/options"
	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/passport_service"
	"os"
	"strconv"
)

func main() {
	options.ParseFlags()
	ps := passport_service.New()

	addr := strconv.FormatInt(options.GlobalConfig.Port, 10)

	s, err := mygrpcadapter.NewMyAuthAdapter(addr, ps)
	if err != nil {
		fmt.Printf("unable to start server: %v", err)
		os.Exit(-1)
	}

	shutdown := make(chan error, 1)
	go func() {
		s.Run(shutdown)
	}()
	_ = <-shutdown
}
