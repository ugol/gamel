package core

type Message struct {
	Body	interface{}
}

func NewMessage() *Message {
	return &Message{}
}