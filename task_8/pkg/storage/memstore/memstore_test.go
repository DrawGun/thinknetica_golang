package memstore

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

func TestDB_StoreDocs_Docs(t *testing.T) {
	store := New()

	err := store.StoreDocs(docs)
	if err != nil {
		t.Errorf("err = %s", err)
		return
	}

	readDocs := store.Docs()
	want := docs[0]
	got := readDocs[0]
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
