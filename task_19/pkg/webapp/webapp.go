// Package webapp запускает web-сервер для доступа к хранилищу
package webapp

import (
	"fmt"
	"net/http"
	"thinknetica_golang/task_19/pkg/structs"

	"text/template"

	"github.com/gorilla/mux"
)

// index представляет собой контракт индекса
type index interface {
	Inspect() map[string][]int
}

// storage представляет собой контракт хранилища
type storage interface {
	Docs() []structs.Document
}

// WebApp структура данных
type WebApp struct {
	ind     index
	store   storage
	address string
}

// New - создает новый экземпляр типа WebApp
func New(ind index, store storage, address string) *WebApp {
	wa := WebApp{
		ind:     ind,
		store:   store,
		address: address,
	}

	return &wa
}

// Run запускает службу для обслуживания запросов
func (wa *WebApp) Run() {
	router := wa.endpoints()
	http.ListenAndServe(wa.address, router)
}

func (wa *WebApp) handleRoot(w http.ResponseWriter, r *http.Request) {
	rootPage := `<h1>Links:</h1><p><a href="/index">Index</a></p><p><a href="/docs">Docs</a></p>`
	_, err := w.Write([]byte(rootPage))
	if err != nil {
		fmt.Println(err)
	}
}

func (wa *WebApp) handleIndex(w http.ResponseWriter, r *http.Request) {
	index := wa.ind.Inspect()
	t := template.New("index")
	t, err := t.Parse(`<p><a href="/"><- back</a></p><h1>Index</h1><ul>{{ range $key, $value := . }}<li>{{ $key }}: {{ $value }}</li>{{ end }}</ul>`)
	if err != nil {
		http.Error(w, "ошибка при обработке шаблона", http.StatusInternalServerError)
		return
	}
	t.Execute(w, index)
}

func (wa *WebApp) handleDocs(w http.ResponseWriter, r *http.Request) {
	index := wa.store.Docs()
	t := template.New("docs")
	t, err := t.Parse(`<p><a href="/"><- back</a></p><h1>Docs</h1><ul>{{ range . }}<li>ID:{{ .ID }} -  {{ .Title }} ({{ .URL }})</li>{{ end }}</ul>`)
	if err != nil {
		http.Error(w, "ошибка при обработке шаблона", http.StatusInternalServerError)
		return
	}
	t.Execute(w, index)
}

func (wa *WebApp) endpoints() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/index", wa.handleIndex).Methods(http.MethodGet)
	router.HandleFunc("/docs", wa.handleDocs).Methods(http.MethodGet)
	router.HandleFunc("/", wa.handleRoot).Methods(http.MethodGet)

	return router
}
