package usecase

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/telegram/bot/user/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/telebot.v4"
)

type UserUsecase interface {
	AskLLM(telebot.Context) error

	ChooseLLM(telebot.Context) error

	ClearConversation(telebot.Context) error
}

type userUsecase struct {
	repo    *repository.UserRepository
	service pb.LLMServiceClient
	logger  *zap.SugaredLogger
}

func New(userRepo *repository.UserRepository, conn *grpc.ClientConn, logger *zap.SugaredLogger) UserUsecase {
	return &userUsecase{
		repo:    userRepo,
		service: pb.NewLLMServiceClient(conn),
		logger:  logger,
	}
}

func (uc *userUsecase) AskLLM(ctx telebot.Context) error {
	id := parseID(ctx)

	uc.logger.Debugw("Getting model from Redis", "id", id)
	model, err := uc.repo.GetModel(id)
	if err != nil {
		model = "deepseek/deepseek-chat-v3-0324:free"
	}

	uc.logger.Debugw("Getting user messages history", "id", id)
	history, _ := uc.repo.GetMessages(id)
	msg := &pb.ChatMessage{
		Role:    "user",
		Content: ctx.Text(),
	}
	history = append(history, msg)

	req := &pb.ChatHistoryRequest{
		User:     &pb.User{Username: ctx.Sender().Username},
		Model:    model,
		Messages: history,
	}

	uc.logger.Debugw("gRPC call to GenerateText", "req", req)
	resp, err := uc.service.GenerateText(context.Background(), req)
	if err != nil {
		uc.logger.Warnw("Generation Error", "err", err)
		return ctx.Send("Ошибка генерации, попробуйте еще раз или смените LLM")
	}

	uc.repo.AddMessage(id, msg)             // user request
	uc.repo.AddMessage(id, &pb.ChatMessage{ // AI answer
		Role:    "assistant",
		Content: resp.Answer,
	})

	return ctx.Send(resp.Answer)
}

func (uc *userUsecase) ClearConversation(ctx telebot.Context) error {
	id := parseID(ctx)

	uc.logger.Debugw("Cleaning conversation", "id", id)
	if err := uc.repo.DeleteMessages(id); err != nil {
		uc.logger.Errorw("Cleaning conversation error", "id", id)
		return ctx.Send("Не удалось очистить историю")
	}

	return ctx.Send("История очищена ✅")
}

func (uc *userUsecase) ChooseLLM(ctx telebot.Context) error {
	id := parseID(ctx)

	model := strings.Replace(ctx.Callback().Data, "select_model|", "", 1)
	model = strings.TrimSpace(model)

	if err := uc.repo.SetModel(id, model); err != nil {
		uc.logger.Errorw("Cant save LLM", "err", err)
		return ctx.Respond(&telebot.CallbackResponse{Text: "Ошибка выбора модели ❌"})
	}

	return ctx.Edit("Модель выбрана ✅")
}

func parseID(ctx telebot.Context) string {
	return strconv.FormatInt(ctx.Sender().ID, 10)
}
