package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("load env.............................")
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to load .env file")
	}
}
