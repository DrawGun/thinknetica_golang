package main

import (
	"pkg/membot"
	"testing"
)

func TestResults(t *testing.T) {
	s := membot.New()
	res, _ := Results(s, "https://yandex.ru/", 2)
	got := len(res)
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
