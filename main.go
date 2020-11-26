package main

import (
	"os"

	"github.com/nekochans/kimono-app-api/infrastructure"
)

func main() {
	region := os.Getenv("REGION")
	if region == "" {
		panic("error setting env REGION")
	}

	userPoolId := os.Getenv("USER_POOL_ID")
	if userPoolId == "" {
		panic("error setting env USER_POOL_ID")
	}

	userPoolClientId := os.Getenv("USER_POOL_WEB_CLIENT_ID")
	if userPoolId == "" {
		panic("error setting env USER_POOL_WEB_CLIENT_ID")
	}

	infrastructure.StartHTTPServer(region, userPoolId, userPoolClientId)
}
