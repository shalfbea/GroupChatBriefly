package chatgpt

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
	"log"
)

const promptStart = "Сделай краткий пересказ истории этой переписки. Оставляй имена:\n"
const testHistory = `Илья: Как здорово прошёл этот день! И погода замечательная в Москве, 18+!
Сергей:Уф уф,а у нас ужасно жарко, +30!
Наташа:А в Брянске + 40,скорее бы вернуться!!
Полина:как же вы много пишете...`

type Chatgpt struct {
	client *openai.Client
	logger *log.Logger
}

func newChatGpt(logger *log.Logger, apiKey string) *Chatgpt {
	return &Chatgpt{
		client: openai.NewClient(apiKey),
		logger: logger,
	}
}
func (chatgpt *Chatgpt) Response(ctx context.Context, chatHistory string) (brief string, err error) {
	content := promptStart + chatHistory
	resp, err := chatgpt.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		chatgpt.logger.Println()
		return
	}

	return resp.Choices[0].Message.Content, nil
}
