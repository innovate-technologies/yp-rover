package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type roverConfig struct {
	ShoutcastKey string `required:"true"`
	RabbitMQURL  string `required:"true" envconfig:"rabbitmq_url"`
}

var config roverConfig

func main() {
	err := envconfig.Process("yprover", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
