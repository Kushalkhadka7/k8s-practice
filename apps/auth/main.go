package main

import (
	"auth/config"
	"auth/router"
	"auth/server"
	"fmt"
	"os"
)

func main() {

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	router := router.New()

	server := server.New(config, router)

	if err := server.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
