package auth

import (
	"strings"
	"testing"
	"thinknetica_golang/task_17/pkg/db"
)

func TestAuth_Сheck(t *testing.T) {
	auth := New()
	user := db.User{
		Login:    "admin",
		Password: "12345",
	}
	tokenString, err := auth.Сheck(user)
	if err != nil {
		t.Errorf("auth.Authorize(); err = %v; want %v", err, nil)
	}

	splitted := strings.Split(tokenString, ".")
	got := len(splitted)
	want := 3

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
