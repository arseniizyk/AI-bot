package main

import (
	"log"
	"os"

	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/grpc/client"
	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/telegram/bot"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := client.Init(os.Getenv("GRPC_PORT"))
	if err != nil {
		logger.Sugar().Errorw("Cant connect to grpc",
			"port", os.Getenv("GRPC_PORT"),
			"err", err,
		)
		log.Fatal(err)
	}
	logger.Info("gRPC connection success")
	defer conn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	cmd := rdb.Ping()
	if cmd.Err() != nil {
		logger.Error("Redis connection error", zap.Error(err))
		log.Fatal(err)
	}

	logger.Info("Redis connection success")

	b := bot.New(os.Getenv("TELEGRAM_API"), logger.Sugar(), conn, rdb)
	if err := b.Init(); err != nil {
		logger.Sugar().Errorw("Cant initialize telegram bot",
			"token", os.Getenv("TELEGRAM_API"),
			"err", err,
		)
	}
	
	b.Run()
}
