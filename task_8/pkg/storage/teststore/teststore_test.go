// Package teststore предоставляет возможность сохранить данные в памяти для тестов
package teststore

import (
	"pkg/crawler"
	"testing"
)

var docs = []crawler.Document{
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

func TestStorage_StoreDocs(t *testing.T) {
	store := New()

	err := store.StoreDocs(docs)
	if err != nil {
		t.Errorf("err = %s", err)
		return
	}

	got := len(store.documents)
	want := 2
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestStorage_Docs(t *testing.T) {
	store := New()
	store.documents = docs

	readDocs := store.Docs()

	got := len(readDocs)
	want := 2
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
