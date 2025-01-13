package main

import (
	"log"

	"github.com/xarick/golang-redis-example/cache"
	"github.com/xarick/golang-redis-example/config"
	"github.com/xarick/golang-redis-example/db"
	"github.com/xarick/golang-redis-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg)

	cache.ConnectRedis(cfg)

	defer db.DB.Close()
	defer cache.RDB.Close()

	r := routes.SetupRouter()
	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
