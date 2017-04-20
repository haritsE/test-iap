package main

import (
	"fmt"

	"github.com/haritsE/test-iap/config"
	"github.com/haritsE/test-iap/router"
	"github.com/kelseyhightower/envconfig"
)

var (
	conf config.Config
)

func configureConfig() {
	err := envconfig.Process("", &conf)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	configureConfig()
	router.CreateRouter(conf)
}
