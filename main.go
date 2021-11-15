package main

import (
	"basic-auth-service/lib"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const defaultPort = "10000"

func main() {
	// load .env file
	err := godotenv.Load()
	if err == nil {
		fmt.Println("loaded config from .env file")
	}

	// load env
	port, hasPort := os.LookupEnv("PORT")
	if !hasPort {
		port = defaultPort
	}

	username, hasUsername := os.LookupEnv("BASIC_AUTH_SERVICE_USERNAME")
	if !hasUsername {
		log.Panic("BASIC_AUTH_SERVICE_USERNAME is mandatory")
	}

	password, hasPassword := os.LookupEnv("BASIC_AUTH_SERVICE_PASSWORD")
	if !hasPassword {
		log.Panic("BASIC_AUTH_SERVICE_PASSWORD is mandatory")
	}

	// http handler and listen
	http.HandleFunc("/", lib.BasicAuthHandler(username, password))

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("listening on %s\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
