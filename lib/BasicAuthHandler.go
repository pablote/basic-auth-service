package lib

import (
	"fmt"
	"net/http"
)

func BasicAuthHandler(expectedUsername, expectedPassword string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req: host=%s, uri=%s\n", r.Host, r.RequestURI)

		username, password, basicAuthOk := r.BasicAuth()
		if !basicAuthOk {
			sendAutheticate(w)
			return
		}

		if username != expectedUsername || password != expectedPassword {
			sendAutheticate(w)
			return
		}

		w.WriteHeader(200)
		_, err := fmt.Fprint(w, "+OK")
		if err != nil {
			fmt.Println(err)
		}
	}
}
