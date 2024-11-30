package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var StartTime = time.Now()

type Config struct {
	PORT string
}

var ENV = LoadEnv()

func LoadEnv() *Config{
	if err:= godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	log.Println("Successfully loaded env file..");
	
	return &Config{
		PORT: getEnv("PORT", ":8080"),
	}

}

func getEnv(key, fallback string) string {
	value, ok:= os.LookupEnv(key)
	
	if !ok {
		return fallback
	}

	return value
}