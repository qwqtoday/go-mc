package msg

import "github.com/Tnze/go-mc/chat"

type EventsHandler struct {
	PlayerChatMessage func(msg chat.Message) error
}