package chat

import (
	"fmt"
	"sync"
)

// Chat реализация сервиса чата
type Chat struct {
	mux      *sync.Mutex
	Clients  []chan string
	MsgQueue chan string
}

// New создает новый объект чата
func New() *Chat {
	c := Chat{
		mux:      &sync.Mutex{},
		Clients:  make([]chan string, 0),
		MsgQueue: make(chan string),
	}

	return &c
}

// Subscribe позволяет участнику подписаться на новые сообщения
func (c *Chat) Subscribe() chan string {
	c.mux.Lock()
	defer c.mux.Unlock()

	client := make(chan string)
	c.Clients = append(c.Clients, client)
	fmt.Println("Subscribe")

	return client
}

// Unsubscribe удаляет участника из подписки
func (c *Chat) Unsubscribe(client chan string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for i := range c.Clients {
		if c.Clients[i] == client {
			c.Clients = append(c.Clients[:i], c.Clients[i+1:]...)
			break
		}
	}
	fmt.Println("Unsubscribe")
}

// Broadcast добавляет сообщение в очередь сообщений
func (c *Chat) Broadcast(message string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.MsgQueue <- message

	fmt.Println("Broadcast")
}
