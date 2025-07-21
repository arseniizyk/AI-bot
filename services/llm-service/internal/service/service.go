package service

import (
	"context"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/arseniizyk/AI-bot/services/llm-service/internal/llm"
	"go.uber.org/zap"
)

type Service struct {
	c *llm.OpenAIClient
	pb.UnimplementedLLMServiceServer
}

func New(c *llm.OpenAIClient) *Service {
	return &Service{
		c: c,
	}
}

func (s *Service) GenerateText(ctx context.Context, req *pb.ChatHistoryRequest) (*pb.TextResponse, error) {
	answer, err := s.c.Ask(req)
	if err != nil {
		zap.L().Error("error while generating response", zap.Error(err))
		return nil, err
	}

	return &pb.TextResponse{
		Answer: answer,
	}, nil
}
