package repository

import (
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/go-redis/redis"
)

type UserRepository struct {
	rdb *redis.Client
}

const maxHistory = 25

func New(rdb *redis.Client) *UserRepository {
	return &UserRepository{
		rdb: rdb,
	}
}

func (r *UserRepository) SetModel(userID string, model string) error {
	key := fmt.Sprintf("user:%s:model", userID)
	return r.rdb.Set(key, model, 0).Err()
}

func (r *UserRepository) GetModel(userID string) (string, error) {
	key := fmt.Sprintf("user:%s:model", userID)
	return r.rdb.Get(key).Result()
}

func (r *UserRepository) AddMessage(userID string, msg *pb.ChatMessage) error {
	key := fmt.Sprintf("user:%s:history", userID)

	data, _ := r.rdb.Get(key).Result()
	var history []*pb.ChatMessage
	if data != "" {
		_ = json.Unmarshal([]byte(data), &history)
	}

	history = append(history, msg)
	if len(history) > maxHistory {
		history = history[len(history)-maxHistory:]
	}

	encoded, err := json.Marshal(history)
	if err != nil {
		return err
	}
	return r.rdb.Set(key, encoded, time.Hour).Err()
}

func (r *UserRepository) DeleteMessages(userID string) error {
	key := fmt.Sprintf("user:%s:history", userID)
	return r.rdb.Del(key).Err()
}

func (r *UserRepository) GetMessages(userID string) ([]*pb.ChatMessage, error) {
	key := fmt.Sprintf("user:%s:history", userID)
	data, err := r.rdb.Get(key).Result()
	if err != nil {
		return nil, err
	}

	var history []*pb.ChatMessage
	err = json.Unmarshal([]byte(data), &history)
	return history, err
}
