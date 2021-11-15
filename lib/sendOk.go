package lib

import (
	"fmt"
	"net/http"
)

func sendOk(w http.ResponseWriter) {
	w.WriteHeader(200)
	_, err := fmt.Fprint(w, "+OK")
	if err != nil {
		fmt.Println(err)
	}
}
