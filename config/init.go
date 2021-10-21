package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(filepath.Join(".env"))
	TimeInit()
}
