package chathistory

import (
	"fmt"
	"strings"

	"github.com/shalfbea/GroupChatBriefly/pkg/logger"
)

type Message struct {
	authorId    int64
	messageList []string
}

func NewMessage(authorId int64, msg string) Message {
	message := Message{
		authorId:    authorId,
		messageList: make([]string, 1),
	}
	message.messageList[0] = msg
	return message
}

type ChatHistory struct {
	authors map[int64]string
	chats   map[int64](*[]Message)
	editables
	logger logger.Logger
}

// AuthorKnown - checks if author is already known and return its name if it is
func (ch *ChatHistory) AuthorKnown(id int64) (string, bool) {
	name, ok := ch.authors[id]
	return name, ok
}

func (ch *ChatHistory) AuthorStore(id int64, name string) {
	ch.authors[id] = name
}

func InitChatHistory(logger logger.Logger) *ChatHistory {
	return &ChatHistory{
		authors: make(map[int64]string),
		chats:   make(map[int64]*[]Message),
		logger:  logger,
	}
}

func addMessage(msgs *[]Message, authorId int64, msg string) (*[]Message, msgIndex) {
	if len(*msgs) > 0 {
		lastMessage := (*msgs)[len(*msgs)-1]
		if authorId == lastMessage.authorId {
			lastMessage.messageList = append(lastMessage.messageList, msg)
			index := len(*msgs) - 1
			(*msgs)[index] = lastMessage
			return msgs, msgIndex{message: len(*msgs), messageIndex: index}
		}
	}
	*msgs = append(*msgs, NewMessage(authorId, msg))
	return msgs, msgIndex{message: len(*&msg), messageIndex: 0}
}

func (ch *ChatHistory) DebugPrintMessages() string {
	res := ""
	for chatId, messages := range ch.chats {
		for _, message := range *messages {
			res += fmt.Sprintf("%d: %d/%s: %s\n", chatId, message.authorId, ch.authors[message.authorId], strings.Join(message.messageList, "."))
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
		res += fmt.Sprintf("%s: %s\n", ch.authors[message.authorId], strings.Join(message.messageList, "."))
	}
	if pop {
		ch.chats[chatId] = emptyMessages()
	}
	return res, nil
}

func (ch *ChatHistory) Store(chatId int64, authorId int64, editable bool, message string) (index int, err error) {
	if _, ok := ch.chats[chatId]; !ok {
		ch.chats[chatId] = emptyMessages()
	}
	var msgIndex msgIndex
	ch.chats[chatId], msgIndex = addMessage(ch.chats[chatId], authorId, message)
	if editable {
		msgIndex.chatId = chatId
		return ch.editables.Store(msgIndex)
	}
	return
}
