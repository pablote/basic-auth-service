package main

import (
	"basic-auth-service/lib"
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "10000"

func main() {
	port, hasPort := os.LookupEnv("PORT")
	if !hasPort {
		port = defaultPort
	}

	username := os.Getenv("BASIC_AUTH_SERVICE_USERNAME")
	password := os.Getenv("BASIC_AUTH_SERVICE_PASSWORD")

	http.HandleFunc("/", lib.BasicAuthHandler(username, password))

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
