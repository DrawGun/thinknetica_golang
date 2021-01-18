package api

import (
	"encoding/json"
	"net/http"
	"pkg/crawler"
	"pkg/engine"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	// IndexedDocuments запрос на индекс
	IndexedDocuments = "/api/v1/index"
	// SEARCH запрос на поиск по соответствию
	SEARCH = "/api/v1/search"
	// DOCUMENTS создание документа, список документов
	DOCUMENTS = "/api/v1/documents"
	// DOCUMENT чтение, обновление и удаление документа
	DOCUMENT = "/api/v1/documents/{docId}"
)

// Service предоставляет интерфейс программного взаимодействия.
type Service struct {
	router *mux.Router
	port   string
	eng    *engine.Service
}

// New - создает новый экземпляр типа Service
func New(eng *engine.Service, port string) *Service {
	s := Service{
		router: mux.NewRouter(),
		eng:    eng,
		port:   port,
	}

	return &s
}

// Run запускает службу для обслуживания запросов
func (s *Service) Run() {
	s.endpoints()
	http.ListenAndServe(s.port, s.router)
}

// Endpoints регистрирует конечные точки Service.
func (s *Service) endpoints() {
	s.router.HandleFunc(IndexedDocuments, s.indexedDocs).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc(DOCUMENTS, s.rawDocs).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc(SEARCH, s.searchDocs).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc(DOCUMENTS, s.createDoc).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc(DOCUMENT, s.getDoc).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc(DOCUMENT, s.updateDoc).Methods(http.MethodPut, http.MethodOptions)
	s.router.HandleFunc(DOCUMENT, s.deleteDoc).Methods(http.MethodDelete, http.MethodOptions)
}

func (s *Service) indexedDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	index := s.eng.Index.Inspect()
	err := json.NewEncoder(w).Encode(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) rawDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	docs := s.eng.Storage.Docs()
	err := json.NewEncoder(w).Encode(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) searchDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := b["query"]
	res, err := s.eng.Search(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) createDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	doc := crawler.Document{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedDocs, err := s.eng.Storage.StoreDocs([]crawler.Document{doc})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.eng.UpdateDocuments(updatedDocs)

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) getDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	docID, err := strconv.Atoi(mux.Vars(r)["docId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	document, ok := s.eng.DocumentByID(docID)
	if ok {
		err := json.NewEncoder(w).Encode(document)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Service) updateDoc(w http.ResponseWriter, r *http.Request) {
	doc := crawler.Document{}
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	docID, err := strconv.Atoi(mux.Vars(r)["docId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doc.ID = docID

	err = s.eng.UpdateDocument(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) deleteDoc(w http.ResponseWriter, r *http.Request) {
	docID, err := strconv.Atoi(mux.Vars(r)["docId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.eng.DeleteDocument(docID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
