package bot

import (
	"strings"
	"time"

	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/telegram/bot/user/repository"
	"github.com/arseniizyk/AI-bot/services/chat-bot/internal/telegram/bot/user/usecase"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/telebot.v4"
)

type Bot struct {
	b      *telebot.Bot
	token  string
	logger *zap.SugaredLogger
	uc     usecase.UserUsecase
}

func New(token string, logger *zap.SugaredLogger, conn *grpc.ClientConn, rdb *redis.Client) *Bot {
	userRepo := repository.New(rdb)

	return &Bot{
		token:  token,
		logger: logger,
		uc:     usecase.New(userRepo, conn, logger),
	}
}

func (b *Bot) Init() error {
	pref := telebot.Settings{
		Token:  b.token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		return err
	}
	b.b = bot

	b.logger.Infow("Bot running", "username", b.b.Me.Username)

	return nil
}

func (b *Bot) Run() {
	b.b.Handle("/start", func(ctx telebot.Context) error {
		ctx.Send("Hello, I'm AI bot, choose what AI you wanna use")
		return SelectKeyboard(ctx)
	})

	b.b.Handle("/ai", func(ctx telebot.Context) error {
		return SelectKeyboard(ctx)
	})

	b.b.Handle("/clear", func(ctx telebot.Context) error {
		return b.uc.ClearConversation(ctx)
	})

	b.b.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return b.uc.AskLLM(ctx)
	})

	b.b.Handle(telebot.OnCallback, func(ctx telebot.Context) error {
		if strings.Contains(ctx.Callback().Data, "select_model") {
			return b.uc.ChooseLLM(ctx)
		}
		return ctx.Respond()
	})

	b.b.Start()
}
