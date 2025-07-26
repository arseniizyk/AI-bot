package formatter

import (
	"regexp"
	"strings"
)

var replacer = strings.NewReplacer(
	"_", "\\_",
	"*", "\\*",
	"[", "\\[",
	"]", "\\]",
	"(", "\\(",
	")", "\\)",
	"~", "\\~",
	"`", "\\`",
	">", "\\>",
	"#", "\\#",
	"+", "\\+",
	"-", "\\-",
	"=", "\\=",
	"|", "\\|",
	"{", "\\{",
	"}", "\\}",
	".", "\\.",
	"!", "\\!",
)

func PreparyForReply(text string) string {
	text = regexp.MustCompile(`(?m)^#{1,6}\s*`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`(?m)^\|.*\|$`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`(?m)^[-*+]\s+`).ReplaceAllString(text, "")
	text = regexp.MustCompile("(?s)```.*?```").ReplaceAllString(text, "")
	return replacer.Replace(text)
}
