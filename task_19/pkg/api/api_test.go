package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"thinknetica_golang/task_19/pkg/engine"
	"thinknetica_golang/task_19/pkg/index/hash"
	"thinknetica_golang/task_19/pkg/search/btree"
	"thinknetica_golang/task_19/pkg/storage/memstore"
	"thinknetica_golang/task_19/pkg/structs"

	"github.com/gorilla/mux"
)

var router *mux.Router

var docs = []structs.Document{
	{
		ID:    0,
		URL:   "https://yandex.ru",
		Title: "Яндекс",
	},
	{
		ID:    1,
		URL:   "https://google.ru",
		Title: "Google",
	},
}

func TestMain(m *testing.M) {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()
	eng := engine.New(store, ind, srch)

	updatedDocs, err := store.StoreDocs(docs)
	if err != nil {
		fmt.Println(err)
		return
	}

	eng.UpdateDocuments(updatedDocs)

	api := New(eng, ":3000")
	api.endpoints()
	router = api.router
	os.Exit(m.Run())
}

func TestService_indexedDocs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, IndexedDocuments, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	got := map[string][]int{}
	_ = json.NewDecoder(rec.Body).Decode(&got)
	want := map[string][]int{
		"":          []int{0, 1},
		"Google":    []int{1},
		"google.ru": []int{1},
		"https":     []int{0, 1},
		"yandex.ru": []int{0},
		"Яндекс":    []int{0},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_rawDocs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, DOCUMENTS, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	got := []structs.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&got)

	want := docs
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_searchDocs(t *testing.T) {
	data := map[string]string{
		"query": "Google",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, SEARCH, bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

	got := []string{}
	_ = json.NewDecoder(rec.Body).Decode(&got)

	want := []string{"https://google.ru - Google"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_createDoc(t *testing.T) {
	data := structs.Document{
		URL:   "example.test",
		Title: "Test Title",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, DOCUMENTS, bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusCreated) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}

}

func TestService_readDoc(t *testing.T) {
	url := strings.Replace(DOCUMENT, "{docId}", "1", 1)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}
	want := structs.Document{ID: 1, URL: "https://google.ru", Title: "Google"}
	got := structs.Document{}
	_ = json.NewDecoder(rec.Body).Decode(&got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_updateDoc(t *testing.T) {
	data := structs.Document{
		URL:   "https://google.ru",
		Title: "Google Test",
	}
	payload, _ := json.Marshal(data)
	url := strings.Replace(DOCUMENT, "{docId}", "1", 1)
	req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}
}

func TestService_deleteDoc(t *testing.T) {
	url := strings.Replace(DOCUMENT, "{docId}", "1", 1)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if !(rec.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rec.Code, http.StatusOK)
	}
}
