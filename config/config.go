package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	wd, _ := os.Getwd()

	curr := wd

	for i := 0; i < 5; i++ {
		envPath := filepath.Join(curr, ".env")
		if _, err := os.Stat(envPath); err == nil {
			// 찾았다면 로드하고 종료
			err := godotenv.Load(envPath)
			if err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}
			log.Printf(".env loaded from: %s", envPath)
			return
		}
		curr = filepath.Dir(curr)
	}
}
