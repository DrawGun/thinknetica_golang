package api

import (
	"fmt"
	"log"
	"net/http"
	"thinknetica_golang/task_18/pkg/chat"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	// SendMessagePath отправка сообщения
	SendMessagePath = "/send"
	// MessagesPath Подписка на сообщения
	MessagesPath = "/messages"
)

var upgrader = websocket.Upgrader{}

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	port   string
	router *mux.Router
	chat   *chat.Chat
}

// New создает объект API
func New(port string, chat *chat.Chat) *API {
	s := API{
		port:   port,
		router: mux.NewRouter(),
		chat:   chat,
	}
	return &s
}

// Run запускает службу для обслуживания запросов
func (api *API) Run() {
	api.endpoints()

	http.ListenAndServe(api.port, api.router)
}

func (api *API) endpoints() {
	api.router.HandleFunc(SendMessagePath, api.handleSendMessage).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc(MessagesPath, api.handleMessages).Methods(http.MethodGet, http.MethodOptions)
}

func (api *API) handleSendMessage(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Enter the password!\n"))
	if err != nil {
		log.Println(err)
	}

	auth := false
	for {
		mt, r, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		resp := string(r)
		if !auth && resp != "password" {
			log.Println("invalid password")

			api.writeMessage(ws, mt, []byte("Invalid password, try again!\n"))
		}

		if auth == true {
			log.Println("Сообщение:", resp)

			api.chat.Broadcast(resp)
			api.writeMessage(ws, mt, []byte("Sent\n"))
		}

		if string(resp) == "password" {
			auth = true
			api.writeMessage(ws, mt, []byte("OK\n"))
		}
	}
}

func (api *API) handleMessages(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	client := api.chat.Subscribe()

	defer ws.Close()
	defer api.chat.Unsubscribe(client)

	// чтение сообщений из канала данного клиента
	for msg := range client {
		err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (api *API) writeMessage(conn *websocket.Conn, mt int, msg []byte) {
	if err := conn.WriteMessage(1, msg); err != nil {
		log.Println(err)
	}
}
