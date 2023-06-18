package telegram

import (
	tele "gopkg.in/telebot.v3"
)

func (b *Bot) StoreHistory(c tele.Context, editable bool, text string) (index int, err error) {
	sender := c.Sender()
	senderName, ok := b.chatHistory.AuthorKnown(sender.ID)
	if !ok {
		senderName = sender.FirstName
		if len(sender.LastName) > 1 {
			senderName += " " + string([]rune(sender.LastName)[0]) + "."
		}
		b.chatHistory.AuthorStore(sender.ID, senderName)
	}
	b.logger.Infof("Store: %s(%d): %s", senderName, sender.ID, text)
	return b.chatHistory.Store(c.Chat().ID, sender.ID, editable, text)
}

func (b *Bot) registerHandlers() {
	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		_, err := b.StoreHistory(c, false, c.Text()) // Расшифровка не требуется, индекс опускаем
		return err
	})

	b.bot.Handle(tele.OnAudio, func(c tele.Context) error {
		_, err := b.StoreHistory(c, false, "*отправил аудио*") //TODO: возможная расшифровка
		return err
	})

	b.bot.Handle(tele.OnMedia, func(c tele.Context) error {
		_, err := b.StoreHistory(c, false, "*отправил медиа*") //TODO: возможная расшифровка
		return err
	})

	b.bot.Handle(tele.OnPhoto, func(c tele.Context) error {
		_, err := b.StoreHistory(c, false, "*отправил фото*") //TODO: возможная расшифровка
		return err
	})
}
