package usecase

import (
	"context"
	"strconv"
	"strings"

	pb "github.com/arseniizyk/AI-bot/proto"
	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/telegram/bot/user/repository"
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
}

func New(userRepo *repository.UserRepository, conn *grpc.ClientConn) UserUsecase {
	return &userUsecase{
		repo:    userRepo,
		service: pb.NewLLMServiceClient(conn),
	}
}

func (uc *userUsecase) AskLLM(ctx telebot.Context) error {
	id := strconv.FormatInt(ctx.Sender().ID, 10)

	model, err := uc.repo.GetModel(id)
	if err != nil {
		model = "deepseek/deepseek-chat-v3-0324:free"
	}

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

	resp, err := uc.service.GenerateText(context.Background(), req)
	if err != nil {
		return ctx.Send("Ошибка генерации: " + err.Error())
	}

	uc.repo.AddMessage(id, msg)
	uc.repo.AddMessage(id, &pb.ChatMessage{
		Role:    "assistant",
		Content: resp.Answer,
	})

	return ctx.Send(resp.Answer)
}

func (uc *userUsecase) ClearConversation(ctx telebot.Context) error {
	id := strconv.FormatInt(ctx.Sender().ID, 10)

	if err := uc.repo.DeleteMessages(id); err != nil {
		return ctx.Send("Не удалось очистить историю: " + err.Error())
	}
	return ctx.Send("История очищена ✅")
}

func (uc *userUsecase) ChooseLLM(ctx telebot.Context) error {
	id := strconv.FormatInt(ctx.Sender().ID, 10)

	model := strings.Replace(ctx.Callback().Data, "select_model|", "", 1)
	model = strings.TrimSpace(model)

	if err := uc.repo.SetModel(id, model); err != nil {
		return ctx.Respond(&telebot.CallbackResponse{Text: "Ошибка выбора модели ❌"})
	}

	return ctx.Edit("Модель выбрана ✅")
}
