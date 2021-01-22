package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// Service представляет собой CLI клиент
type Service struct {
	reader *bufio.Reader
}

func new() *Service {
	s := Service{
		reader: bufio.NewReader(os.Stdin),
	}

	return &s
}

func main() {
	service := new()
	go service.subscribe()
	service.dialog()
}

func (s *Service) dialog() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8888/send", nil)
	if err != nil {
		log.Fatalf("не удалось подключиться к серверу: %v", err)
		conn.Close()
	}
	defer conn.Close()

	_, resp, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("не удалось прочитать сообщение: %v", err)
	}
	fmt.Print(string(resp))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = strings.ReplaceAll(text, "\r\n", "")
		text = strings.ReplaceAll(text, "\n", "")

		if strings.Compare("", text) == 0 || strings.Compare("exit", text) == 0 {
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			conn.Close()
			log.Fatalf("не удалось отправить сообщение: %v", err)
		}

		_, resp, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("не удалось прочитать сообщение: %v", err)
		}
		fmt.Printf("Сообщеиние от сервера: %v\n", string(resp))
	}
}

func (s *Service) subscribe() {
	conn, r, err := websocket.DefaultDialer.Dial("ws://localhost:8888/messages", nil)
	if err != nil {
		fmt.Println(err, r.StatusCode)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			fmt.Println(err)
			return
		}

		fmt.Printf("Сообщение подписчикам: %v\n", string(message))
	}
}
