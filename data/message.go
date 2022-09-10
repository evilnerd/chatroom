package data

import "time"

type ReceivedMessage struct {
	Timestamp time.Time `json:"timestamp"`
	Msg       string    `json:"msg"`
	Sender    User      `json:"sender"`
}

type SendMessage struct {
	Body string `json:"body"`
}

func NewReceivedMessage(sender User, msg string) ReceivedMessage {
	return ReceivedMessage{
		Timestamp: time.Now(),
		Msg:       msg,
		Sender:    sender,
	}
}

func NewSendMessage(body string) SendMessage {
	return SendMessage{
		body,
	}
}
