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
	CLIENT_ID string
	CLIENT_SECRET string
	CALLBACK_URL string
}

var Env = LoadEnv()

func LoadEnv() *Config{
	if err:= godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	log.Println("Successfully loaded env file..");
	
	return &Config{
		PORT: getEnv("PORT", ":8080"),
		CLIENT_ID: getEnv("CLIENT_ID", "<client-id>"),
		CLIENT_SECRET: getEnv("CLIENT_SECRET", "<client-secret>"),
		CALLBACK_URL: getEnv("CALLBACK_URL", "<callback-url>"),
	}

}

func getEnv(key, fallback string) string {
	value, ok:= os.LookupEnv(key)
	
	if !ok {
		return fallback
	}

	return value
}