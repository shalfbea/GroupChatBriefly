package telegram

import (
	"time"

	"github.com/shalfbea/GroupChatBriefly/pkg/chatgpt"
	"github.com/shalfbea/GroupChatBriefly/pkg/config"
	"github.com/shalfbea/GroupChatBriefly/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	bot     *tele.Bot
	config  *config.Config
	logger  logger.Logger
	chatgpt *chatgpt.Chatgpt
}

func NewBot(config *config.Config, logger logger.Logger, chatgpt *chatgpt.Chatgpt) (bot *Bot, err error) {
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
		bot:     telebot,
		config:  config,
		logger:  logger,
		chatgpt: chatgpt,
	}

	//Just for testing. Will be moved to smth like "register handlers"
	telebot.Handle("/start", func(c tele.Context) error {
		return c.Send(config.Messages.Start)
	})

	return bot, err
}

func RunBot(b *Bot) {
	b.logger.Info("Bot started!")
	b.bot.Start()
}
