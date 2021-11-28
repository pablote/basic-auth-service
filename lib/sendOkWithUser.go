package lib

import (
	"fmt"
	"net/http"
)

func sendOkWithUser(w http.ResponseWriter, username string) {
	w.Header().Set("x-auth-username", username)
	w.WriteHeader(200)
	_, err := fmt.Fprint(w, "+OK")
	if err != nil {
		fmt.Println(err)
	}
}
