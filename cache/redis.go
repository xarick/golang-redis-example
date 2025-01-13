package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xarick/golang-redis-example/config"
)

var (
	RDB *redis.Client          // Redis mijoz obyekti
	Ctx = context.Background() // Kontekst Redis uchun
)

// ConnectRedis Redis serveriga ulanish funksiyasi
func ConnectRedis(cfg config.Application) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr, // Redis serverining manzili va porti
		Password: cfg.RedisPass, // Parol (agar mavjud bo'lmasa, bo'sh qoldiriladi)
		DB:       cfg.RedisDB,   // Redisning ma'lumotlar bazasi (standart 0)
	})

	// Ulanishni tekshirish
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

// SetCache ma'lumotni Redisga saqlash
func SetCache(key string, value interface{}, expiration time.Duration) error {
	err := RDB.Set(Ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Failed to set cache for key %s: %v", key, err)
	}
	return err
}

// GetCache Redisdan ma'lumotni olish
func GetCache(key string) (string, error) {
	val, err := RDB.Get(Ctx, key).Result()
	if err == redis.Nil {
		log.Printf("Key %s does not exist", key)
		return "", nil
	} else if err != nil {
		log.Printf("Failed to get cache for key %s: %v", key, err)
		return "", err
	}
	return val, nil
}

// IncrementCache ma'lumotni oshirish
func IncrementCache(key string) (int64, error) {
	val, err := RDB.Incr(Ctx, key).Result()
	if err != nil {
		log.Printf("Failed to increment cache for key %s: %v", key, err)
	}
	return val, err
}

// DeleteCache ma'lumotni Redisdan o'chirish
func DeleteCache(key string) error {
	err := RDB.Del(Ctx, key).Err()
	if err != nil {
		log.Printf("Failed to delete cache for key %s: %v", key, err)
	}
	return err
}

// Redis buyruq          | Tavsif                               | Golang usuli
// ----------------------|--------------------------------------|----------------------------------------
// SET key value         | Kalitga qiymat yozish               | Set(ctx, key, value, ttl)
// GET key               | Kalitning qiymatini olish           | Get(ctx, key)
// DEL key               | Kalitni o'chirish                   | Del(ctx, key)
// EXISTS key            | Kalit mavjudligini tekshirish       | Exists(ctx, key)
// INCR key              | Kalitning qiymatini 1 ga oshirish   | Incr(ctx, key)
// HSET hash field value | Hash maydonini qiymat bilan o'rnatish| HSet(ctx, hash, field, value)
// HGET hash field       | Hash maydonidan qiymat olish        | HGet(ctx, hash, field)
// LPUSH key value       | List boshiga qiymat qo'shish        | LPush(ctx, key, value)
// LRANGE key start stop | Listdagi elementlarni olish         | LRange(ctx, key, start, stop)
// SADD key value        | Set-ga qiymat qo'shish              | SAdd(ctx, key, value)
// SMEMBERS key          | Set-dagi barcha qiymatlarni olish   | SMembers(ctx, key)
