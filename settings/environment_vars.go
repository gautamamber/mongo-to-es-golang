package settings

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironmnetVars() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
