package main

import (
	"flightlesssomething"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var (
	version string
)

func main() {
	godotenv.Load()
	c, err := flightlesssomething.NewConfig()
	if err != nil {
		log.Fatalln("Failed to get config:", err)
	}

	if c.Version {
		fmt.Println("Version:", version)
		return
	}

	flightlesssomething.Start(c, version)
}
