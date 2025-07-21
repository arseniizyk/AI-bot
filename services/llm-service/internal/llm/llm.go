package llm

import (
	"context"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type OpenAIClient struct {
	c      *openai.Client
	logger *zap.SugaredLogger
}

func New(c *openai.Client, logger *zap.SugaredLogger) *OpenAIClient {
	return &OpenAIClient{
		c:      c,
		logger: logger,
	}
}

func (l *OpenAIClient) Ask(req *pb.ChatHistoryRequest) (string, error) {
	var messages []openai.ChatCompletionMessage

	for _, m := range req.Messages {
		role := openai.ChatMessageRoleUser
		if m.Role == "assistant" {
			role = openai.ChatMessageRoleAssistant
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: m.Content,
		})
	}

	resp, err := l.c.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    req.Model,
		Messages: messages,
	})
	if err != nil {
		l.logger.Errorw("ChatCompletion error",
			"messages", messages,
			"error", err,
			"model", req.Model,
		)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
