package telegram

import (
	"context"
	"github.com/hashicorp/go-multierror"
	tele "gopkg.in/telebot.v3"
)

// registerCommands register handlers for bot commands
func (b *Bot) registerCommands() {
	b.bot.Handle("/start", func(c tele.Context) error {
		return c.Send(b.config.Messages.Start)
	})

	b.bot.Handle("/preview", func(c tele.Context) error {
		history, err := b.chatHistory.PopMessagesForChat(c.Chat().ID, false)
		if err != nil {
			c.Send(err.Error())
		}
		return c.Send(history)
	})

	b.bot.Handle("/brief", func(c tele.Context) error {
		msg, err := b.bot.Send(c.Chat(), b.config.Messages.LoadingBrief)
		if err != nil {
			return err
		}
		history, err := b.chatHistory.PopMessagesForChat(c.Chat().ID, true)
		if err != nil {
			_, errEdit := b.bot.Edit(msg, err.Error())
			return multierror.Append(err, errEdit)
		}
		briefly, err := b.chatgpt.Response(context.Background(), history)
		if err != nil {
			_, errEdit := b.bot.Edit(msg, err.Error())
			return multierror.Append(err, errEdit)
		}
		_, err = b.bot.Edit(msg, briefly)
		return err
	})

	b.setCommands()
}

// setCommands set's available bot commands text and description
func (b *Bot) setCommands() {
	commands := []tele.Command{
		{
			Text:        "/preview",
			Description: "Превью сохраненного диалога",
		},
		{
			Text:        "/brief",
			Description: "Превью сохраненного ",
		},
	}
	b.bot.SetCommands(commands)
}
