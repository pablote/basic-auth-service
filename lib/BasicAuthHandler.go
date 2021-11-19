package lib

import (
	"fmt"
	"github.com/gobwas/glob"
	htpasswd2 "github.com/tg123/go-htpasswd"
	"log"
	"net/http"
	"strings"
)

func BasicAuthHandler(expectedUsername, expectedPassword, htpasswd string, hostAllowList, pathAllowList []string) http.HandlerFunc {
	parsedHostAllowList := make([]glob.Glob, 0, len(hostAllowList))
	for _, host := range hostAllowList {
		parsedHostAllowList = append(parsedHostAllowList, glob.MustCompile(host))
	}

	parsedPathAllowList := make([]glob.Glob, 0, len(pathAllowList))
	for _, path := range pathAllowList {
		parsedPathAllowList = append(parsedPathAllowList, glob.MustCompile(path))
	}

	var htpasswdFile *htpasswd2.File
	var err error
	if len(htpasswd) > 0 {
		htpasswdFile, err = htpasswd2.NewFromReader(strings.NewReader(htpasswd), htpasswd2.DefaultSystems, nil)
		if err != nil {
			log.Fatal("failed to parse htpasswd")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// check if in allow list, if not allow request
		matchesHost := false
		for _, hostGlob := range parsedHostAllowList {
			if hostGlob.Match(r.Host) {
				matchesHost = true
				break
			}
		}

		matchesPath := false
		for _, pathGlob := range parsedPathAllowList {
			if pathGlob.Match(r.RequestURI) {
				matchesPath = true
				break
			}
		}

		fmt.Printf("req: host=%s uri=%s matchesHost=%t matchesPath=%t\n", r.Host, r.RequestURI, matchesHost, matchesPath)

		if !matchesHost || !matchesPath {
			sendOk(w)
			return
		}

		// return "WWW-Authenticate" if no auth header or wrong username/password
		username, password, basicAuthOk := r.BasicAuth()
		if !basicAuthOk {
			sendAutheticate(w)
			return
		}

		// authenticate either using htpasswd or username/password
		if htpasswdFile != nil {
			if !htpasswdFile.Match(username, password) {
				sendAutheticate(w)
				return
			}
		} else {
			if username != expectedUsername || password != expectedPassword {
				sendAutheticate(w)
				return
			}
		}

		// allow request
		sendOk(w)
	}
}
