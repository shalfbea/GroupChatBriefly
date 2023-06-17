package chathistory

import "errors"

// Нужен чтобы затем присваивать по этим индексам новое сообщение
type msgIndex struct {
	chatId       int64
	message      int
	messageIndex int
}

const maxEditables = 100

var (
	errEditablesLimit    = errors.New("Can't store more editables")
	errEditablesNotFound = errors.New("Can't find that editable")
)

type editables []*msgIndex

func (i *editables) Init() editables {
	return make([]*msgIndex, 0, maxEditables)
}

func (e editables) Store(index msgIndex) (int, error) {
	for key, v := range e {
		if v == nil {
			v = &index
			return key, nil
		}
	}
	if len(e) < maxEditables {
		e = append(e, &index)
		return len(e) - 1, nil
	}
	return 0, errEditablesLimit
}

func (e editables) Pop(index int) (msgIndex msgIndex, err error) {
	if len(e) < index {
		return msgIndex, errEditablesNotFound
	}
	if e[index] == nil {
		return msgIndex, errEditablesNotFound
	}
	return *(e[index]), nil
}
