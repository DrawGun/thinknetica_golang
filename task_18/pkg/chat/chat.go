package chat

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Chat реализация сервиса чата
type Chat struct {
	mux      *sync.Mutex
	Clients  map[string]chan string
	MsgQueue chan string
}

// New создает новый объект чата
func New() *Chat {
	c := Chat{
		mux:      &sync.Mutex{},
		Clients:  make(map[string]chan string, 0),
		MsgQueue: make(chan string),
	}

	return &c
}

// Subscribe позволяет участнику подписаться на новые сообщения
func (c *Chat) Subscribe() (string, chan string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 32
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	strID := b.String()
	client := make(chan string)
	c.Clients[strID] = client

	return strID, client
}

// Unsubscribe удаляет участника из подписки
func (c *Chat) Unsubscribe(cID string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.Clients, cID)
}

// Broadcast добавляет сообщение в очередь сообщений
func (c *Chat) Broadcast(message string) {
	c.MsgQueue <- message
}
