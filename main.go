package main

import (
	"basic-auth-service/config"
	"basic-auth-service/lib"
	"fmt"
	"net/http"
)

func main() {
	conf := config.New()

	http.HandleFunc("/", lib.BasicAuthHandler(conf.Username, conf.Password, conf.HostAllowList, conf.PathAllowList))

	addr := fmt.Sprintf("0.0.0.0:%s", conf.Port)
	fmt.Printf("listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
