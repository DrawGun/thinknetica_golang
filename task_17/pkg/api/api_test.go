package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"thinknetica_golang/task_17/pkg/db"
)

var api *API

func TestMain(m *testing.M) {
	api = New(APIPort)
	api.endpoints()

	os.Exit(m.Run())
}

func TestAPI_handleAuth(t *testing.T) {
	user := db.User{
		Login:    "admin",
		Password: "12345",
	}
	payload, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, Auth, bytes.NewReader(payload))
	rec := httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}

	token := rec.Body.String()
	splitted := strings.Split(token, ".")
	got := len(splitted)
	want := 3

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	user = db.User{
		Login:    "admin",
		Password: "11111",
	}
	payload, _ = json.Marshal(user)
	req = httptest.NewRequest(http.MethodPost, Auth, bytes.NewReader(payload))
	rec = httptest.NewRecorder()
	api.router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Status code = %d; want %d", rec.Code, http.StatusOK)
	}
}
