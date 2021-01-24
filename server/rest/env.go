package rest

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// init config
	if err := godotenv.Load(); err != nil {
		log.Println("NO .env file found on server")
	}
}

func RESTVersion() string {
	return os.Getenv("APP_VERSION")
}

func RESTServerName() string {
	return os.Getenv("APP_NAME")
}

func RESTDataPath() string {
	return os.Getenv("APP_DATA_PATH")
}

func RESTDataFileName() string {
	return os.Getenv("APP_DATA_FILENAME")
}
