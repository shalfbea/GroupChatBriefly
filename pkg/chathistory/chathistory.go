package chathistory

import (
	"fmt"

	"github.com/shalfbea/GroupChatBriefly/pkg/logger"
)

type Author struct {
	id   int64
	name string
}

type Message struct {
	author  Author
	message string
}

type ChatHistory struct {
	chats  map[int64](*[]Message)
	logger logger.Logger
}

func InitChatHistory(logger logger.Logger) *ChatHistory {
	return &ChatHistory{
		chats:  make(map[int64]*[]Message),
		logger: logger,
	}
}

func addMessage(msgs *[]Message, authorId int64, authorName, msg string) *[]Message {
	if len(*msgs) > 0 {
		lastMessage := (*msgs)[len(*msgs)-1]
		if authorId == lastMessage.author.id {
			lastMessage.message = lastMessage.message + "." + msg
			(*msgs)[len(*msgs)-1] = lastMessage
			return msgs
		}
	}
	*msgs = append(*msgs, Message{author: Author{authorId, authorName}, message: msg})
	return msgs
}

func (ch *ChatHistory) DebugPrintMessages() string {
	res := ""
	for chatId, messages := range ch.chats {
		for _, message := range *messages {
			res += fmt.Sprintf("%d: %d/%s: %s\n", chatId, message.author.id, message.author.name, message.message)
		}
	}
	return res
}

func emptyMessages() *[]Message {
	tmp := make([]Message, 0)
	return &tmp
}

// PopMessagesForChat pops history for chat and also deletes history
func (ch *ChatHistory) PopMessagesForChat(chatId int64, pop bool) (string, error) {
	if _, ok := ch.chats[chatId]; !ok {
		return "", fmt.Errorf("no info found for chat %d", chatId)
	}
	if len(*ch.chats[chatId]) == 0 {
		return "", fmt.Errorf("no info found for chat %d", chatId)
	}
	res := ""
	for _, message := range *ch.chats[chatId] {
		res += fmt.Sprintf("%s: %s\n", message.author.name, message.message)
	}
	if pop {
		ch.chats[chatId] = emptyMessages() //for testing not deleting
	}
	return res, nil
}

func (ch *ChatHistory) Store(chatId int64, authorId int64, authorName string, message string) {
	if _, ok := ch.chats[chatId]; !ok {
		ch.chats[chatId] = emptyMessages()
	}
	ch.chats[chatId] = addMessage(ch.chats[chatId], authorId, authorName, message)
}
