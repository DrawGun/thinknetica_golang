package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"thinknetica_golang/task_17/pkg/auth"
	"thinknetica_golang/task_17/pkg/db"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Auth запрос на проверку учетной записи
const Auth = "/auth"

// APIPort константа для указания порта
const APIPort = ":8888"

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	port   string
	router *mux.Router
}

// New создает объект API
func New(port string) *API {
	s := API{
		port:   port,
		router: mux.NewRouter(),
	}
	return &s
}

// Run запускает службу для обслуживания запросов
func (api *API) Run() {
	api.endpoints()

	loggedRouter := handlers.LoggingHandler(os.Stdout, api.router)
	http.ListenAndServe(api.port, loggedRouter)
}

func (api *API) endpoints() {
	api.router.HandleFunc(Auth, api.handleAuth).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) handleAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow", "OPTIONS, POST")

	user := db.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth := auth.New()

	token, err := auth.Сheck(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		fmt.Println(err)
	}
}
