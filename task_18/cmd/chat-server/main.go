package main

import (
	"thinknetica_golang/task_18/pkg/api"
	"thinknetica_golang/task_18/pkg/chat"
)

// APIPort константа для указания порта
const APIPort = ":8888"

// Service представляет собой чат-сервер
type Service struct {
	api  *api.API
	chat *chat.Chat
}

func main() {
	service := new()
	go service.publishMessages()
	service.api.Run()
}

func new() *Service {
	chat := chat.New()

	s := Service{
		api:  api.New(APIPort, chat),
		chat: chat,
	}

	return &s
}

func (s *Service) publishMessages() {
	for msg := range s.chat.MsgQueue {
		for _, c := range s.chat.Clients {
			c <- msg
		}
	}
}
