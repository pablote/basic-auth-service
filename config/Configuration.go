package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

const defaultPort = "10000"
const prefix = "BASIC_AUTH_SERVICE"

type Configuration struct {
	Port          string
	Username      string
	Password      string
	HostAllowList []string
	PathAllowList []string
	HtPasswd      string
}

func New() Configuration {
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

	username := os.Getenv(fmt.Sprintf("%s_USERNAME", prefix))
	password := os.Getenv(fmt.Sprintf("%s_PASSWORD", prefix))
	htPasswd := os.Getenv(fmt.Sprintf("%s_HTPASSWD", prefix))
	if !((len(username) > 0 && len(password) > 0) || len(htPasswd) > 0 ){
		log.Fatal("either USERNAME and PASSWORD or HTPASSWD needs to be defined")
	}

	hostAllowList, hasHostAllowList := os.LookupEnv(fmt.Sprintf("%s_HOST_ALLOWLIST", prefix))
	if !hasHostAllowList {
		hostAllowList = "*"
	}

	hostAllowListAsList := strings.Split(hostAllowList, ",")
	for i, host := range hostAllowListAsList {
		hostAllowListAsList[i] = strings.TrimSpace(host)
	}

	pathAllowList, hasPathAllowList := os.LookupEnv(fmt.Sprintf("%s_PATH_ALLOWLIST", prefix))
	if !hasPathAllowList {
		pathAllowList = "*"
	}

	pathAllowListAsList := strings.Split(pathAllowList, ",")
	for i, path := range pathAllowListAsList {
		pathAllowListAsList[i] = strings.TrimSpace(path)
	}

	config := Configuration{
		Port:          port,
		Username:      username,
		Password:      password,
		HtPasswd:      htPasswd,
		HostAllowList: hostAllowListAsList,
		PathAllowList: pathAllowListAsList,
	}
	fmt.Printf("config: %+v\n", config)

	return config
}
