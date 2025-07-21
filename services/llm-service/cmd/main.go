package main

import (
	"log"
	"net"
	"os"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/arseniizyk/AI-bot/services/llm-service/internal/llm"
	"github.com/arseniizyk/AI-bot/services/llm-service/internal/service"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Logger started")

	cfg := openai.DefaultConfig(os.Getenv("OPENROUTER_API"))
	cfg.BaseURL = "https://openrouter.ai/api/v1"

	llmClient := llm.New(openai.NewClientWithConfig(cfg), logger.Sugar())
	grpcServer := grpc.NewServer()
	pb.RegisterLLMServiceServer(grpcServer, service.New(llmClient))

	logger.Info("gRPC LLM Service registered")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
