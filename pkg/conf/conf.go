package conf

import (
	"os"
	"sync"

	godotenv "github.com/joho/godotenv"
)

var (
	config *ConfigStruct
	once   sync.Once
)

func InitConfig() {
	godotenv.Load("./.env")

	config = &ConfigStruct{
		BotToken:         os.Getenv("BOT_TOKEN"),
		DBPath:           DBPath,
		DBMigrationsPath: DBMigrationsPath,
	}
}

func GetConfig() *ConfigStruct {
	once.Do(InitConfig)

	return config
}

type ConfigStruct struct {
	BotToken         string `env:"BOT_TOKEN"`
	DBPath           string
	DBMigrationsPath string
}
