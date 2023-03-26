package require

import (
	"log"
	"os"
)

const APITokenEnv = "API_TOKEN"

func APIToken() string {
	token := os.Getenv(APITokenEnv)
	if len(token) == 0 {
		log.Fatal("missing api token")
	}

	return token
}
