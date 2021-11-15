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
	Port string
	Username string
	Password string
	HostAllowList []string
	PathAllowList []string
}

func New () Configuration {
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

	username, hasUsername := os.LookupEnv(fmt.Sprintf("%s_USERNAME", prefix))
	if !hasUsername {
		log.Panic(fmt.Sprintf("%s_USERNAME is mandatory", prefix))
	}

	password, hasPassword := os.LookupEnv(fmt.Sprintf("%s_PASSWORD", prefix))
	if !hasPassword {
		log.Panic(fmt.Sprintf("%s_PASSWORD is mandatory", prefix))
	}

	hostAllowList, hasHostAllowList := os.LookupEnv(fmt.Sprintf("%s_HOST_ALLOWLIST", prefix))
	if !hasHostAllowList {
		hostAllowList = "*"
	}

	hostAllowListAsList := strings.Split(hostAllowList,",")
	for i, host := range hostAllowListAsList {
		hostAllowListAsList[i] = strings.TrimSpace(host)
	}

	pathAllowList, hasPathAllowList := os.LookupEnv(fmt.Sprintf("%s_PATH_ALLOWLIST", prefix))
	if !hasPathAllowList {
		pathAllowList = "*"
	}

	pathAllowListAsList := strings.Split(pathAllowList,",")
	for i, path := range pathAllowListAsList {
		pathAllowListAsList[i] = strings.TrimSpace(path)
	}

	return Configuration{
		Port: port,
		Username: username,
		Password: password,
		HostAllowList: hostAllowListAsList,
		PathAllowList: pathAllowListAsList,
	}
}
