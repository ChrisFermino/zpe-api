package utils

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func EnvVariable(key string) string {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("error loading .env file")
    }
    return os.Getenv(key)
}
