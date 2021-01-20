package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"thinknetica_golang/task_18/pkg/api"

	"github.com/gorilla/websocket"
)

// APIPort константа для указания порта
const APIPort = ":8888"

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
	url := "ws://localhost" + APIPort + api.SendMessagePath
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
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
		msg, _ := reader.ReadString('\n')
		msg = strings.Replace(msg, "\n", "", -1)

		if msg == "" {
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
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
	url := "ws://localhost" + APIPort + api.MessagesPath
	conn, r, err := websocket.DefaultDialer.Dial(url, nil)
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
