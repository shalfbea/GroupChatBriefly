package telegram

import (
	"context"
	"time"

	"github.com/shalfbea/GroupChatBriefly/pkg/chatgpt"
	"github.com/shalfbea/GroupChatBriefly/pkg/chathistory"
	"github.com/shalfbea/GroupChatBriefly/pkg/config"
	"github.com/shalfbea/GroupChatBriefly/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot         *tele.Bot
	config      *config.Config
	logger      logger.Logger
	chatgpt     *chatgpt.Chatgpt
	chathistory *chathistory.ChatHistory
}

func NewBot(config *config.Config, logger logger.Logger, chatgpt *chatgpt.Chatgpt, chatHistory *chathistory.ChatHistory) (bot *Bot, err error) {
	prefs := tele.Settings{
		Token:  config.TelegramToken,
		Poller: &tele.LongPoller{Timeout: time.Duration(config.PollingTimeout) * time.Second},
	}
	telebot, err := tele.NewBot(prefs)
	if err != nil {
		logger.Errorf("token : %s, Problem with newBot: %v", config.TelegramToken, err)
		return nil, err
	}
	bot = &Bot{
		bot:         telebot,
		config:      config,
		logger:      logger,
		chatgpt:     chatgpt,
		chathistory: chatHistory,
	}
	//Just for testing. Will be moved to smth like "register handlers"
	telebot.Handle("/start", func(c tele.Context) error {
		return c.Send(config.Messages.Start)
	})

	//TODO : refactor
	telebot.Handle(tele.OnText, func(c tele.Context) error {
		sender := c.Sender()
		senderName := sender.FirstName
		if len(sender.LastName) > 1 {
			senderName += " " + string([]rune(sender.LastName)[0]) + "."
		}
		logger.Info("Recv: %v, %v, %v", c.Chat(), senderName, c.Text())
		chatHistory.Store(c.Chat().ID, c.Sender().ID, senderName, c.Message().Text)
		return nil
	})

	telebot.Handle("/preview", func(c tele.Context) error {
		history, err := chatHistory.PopMessagesForChat(c.Chat().ID, false)
		if err != nil {
			c.Send(err.Error())
		}
		return c.Send(history)
	})

	telebot.Handle("/brief", func(c tele.Context) error {
		history, err := chatHistory.PopMessagesForChat(c.Chat().ID, true)
		if err != nil {
			c.Send(err.Error())
		}
		briefly, err := chatgpt.Response(context.Background(), history)
		if err != nil {
			c.Send(err.Error())
		}
		return c.Send(briefly)
	})
	return bot, err
}

func RunBot(b *Bot) {
	b.logger.Info("Bot started!")
	b.bot.Start()
}
