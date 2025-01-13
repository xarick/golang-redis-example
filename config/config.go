package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string

	PgHost   string
	PgPort   string
	PgUser   string
	PgPass   string
	PgDbName string

	RedisAddr string
	RedisPass string
	RedisDB   int
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")

	cfg.PgHost = os.Getenv("PG_HOST")
	cfg.PgPort = os.Getenv("PG_PORT")
	cfg.PgUser = os.Getenv("PG_USER")
	cfg.PgPass = os.Getenv("PG_PASS")
	cfg.PgDbName = os.Getenv("PG_DB_NAME")

	cfg.RedisAddr = os.Getenv("REDIS_ADDR")
	cfg.RedisPass = os.Getenv("REDIS_PASS")

	RedisDBStr := os.Getenv("REDIS_DB")
	cfg.RedisDB, err = strconv.Atoi(RedisDBStr)
	if err != nil {
		fmt.Println("REDIS_DB aniqlanmagan, standart qiymat: 0")
		cfg.RedisDB = 0
	}

	return cfg
}
