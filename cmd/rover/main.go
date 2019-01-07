package main

import (
	"fmt"
	"log"
	"os"

	configpkg "github.com/innovate-technologies/yp-rover/internal/config"
	"github.com/kelseyhightower/envconfig"
)

var config configpkg.Config

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
