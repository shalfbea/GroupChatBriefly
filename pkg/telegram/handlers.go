package telegram

import (
	"context"
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) registerHandlers() {
	b.bot.Handle("/start", func(c tele.Context) error {
		return c.Send(b.config.Messages.Start)
	})

	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		sender := c.Sender()
		senderName := sender.FirstName
		if len(sender.LastName) > 1 {
			senderName += " " + string([]rune(sender.LastName)[0]) + "."
		}
		b.logger.Info("Recv: %v, %v, %v", c.Chat(), senderName, c.Text())
		b.chatHistory.Store(c.Chat().ID, c.Sender().ID, senderName, c.Message().Text)
		return nil
	})

	b.bot.Handle("/preview", func(c tele.Context) error {
		history, err := b.chatHistory.PopMessagesForChat(c.Chat().ID, false)
		if err != nil {
			c.Send(err.Error())
		}
		return c.Send(history)
	})

	b.bot.Handle("/brief", func(c tele.Context) error {
		history, err := b.chatHistory.PopMessagesForChat(c.Chat().ID, true)
		if err != nil {
			c.Send(err.Error())
		}
		briefly, err := b.chatgpt.Response(context.Background(), history)
		if err != nil {
			c.Send(err.Error())
		}
		return c.Send(briefly)
	})
}
