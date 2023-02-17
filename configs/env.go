package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func EnvDbHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DBHOST")
}

func EnvDbName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DBNAME")
}

func EnvDbUser() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DBUSER")
}

func EnvDbPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DBPASSWORD")
}

func EnvDbPort() int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	i, _ := strconv.Atoi(os.Getenv("DBPORT"))
	return i
}

func EnvJwtSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("JWTSECRET")
}

func EnvJwtExpirySeconds() int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	i, _ := strconv.Atoi(os.Getenv("JWTEXPIRYINSECONDS"))
	return i
}
