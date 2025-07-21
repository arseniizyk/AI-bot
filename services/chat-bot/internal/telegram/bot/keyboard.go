package bot

import "gopkg.in/telebot.v4"

var models = map[string]string{
	"DeepSeek-V3":           "deepseek/deepseek-chat-v3-0324:free",
	"Qwen QwQ 32B":          "qwen/qwq-32b:free",
	"Qwen 2.5 72B Instruct": "qwen/qwen-2.5-72b-instruct:free",
	"MetaLlama3370B":        "meta-llama/llama-3.3-70b-instruct:free",
}

func SelectKeyboard(ctx telebot.Context) error {
	inline := &telebot.ReplyMarkup{}

	var buttons []telebot.Row
	for name, value := range models {
		btn := inline.Data(name, "select_model", value)
		buttons = append(buttons, inline.Row(btn))
	}

	inline.Inline(buttons...)
	return ctx.Send("Выбери модель:", inline)
}
