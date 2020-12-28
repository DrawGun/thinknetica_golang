// Package webapp запускает web-сервер для доступа к хранилищу
package webapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pkg/index/hash"
	"pkg/storage/memstore"
	"testing"

	"github.com/gorilla/mux"
)

func TestWebApp_Run(t *testing.T) {
	store := memstore.New()
	ind := hash.New()
	wa := New(ind, store, ":8000")
	router := wa.endpoints()

	tests := []struct {
		name  string
		route string
	}{
		{name: "Root", route: "/"},
		{name: "Index", route: "/index"},
		{name: "Docs", route: "/docs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := []int{}
			payload, _ := json.Marshal(data)
			req := httptest.NewRequest(http.MethodPost, tt.route, bytes.NewBuffer(payload))
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			if rr.Code != http.StatusMethodNotAllowed {
				t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
			}

			req = httptest.NewRequest(http.MethodGet, tt.route, nil)
			rr = httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			if rr.Code != http.StatusOK {
				t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
			}
		})
	}
}

func TestWebApp_HandleRoot(t *testing.T) {
	store := memstore.New()
	ind := hash.New()
	wa := New(ind, store, ":8000")
	mux := mux.NewRouter()
	mux.HandleFunc("/", wa.handleRoot).Methods("GET")

	data := []int{}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rr = httptest.NewRecorder()

	mux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}
