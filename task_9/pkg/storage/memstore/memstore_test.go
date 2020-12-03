package memstore

import (
	"pkg/crawler"
	"pkg/teststore"
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
		t.Fatalf("err = %q", err)
	}

	readDocs := store.Docs()
	want := docs[0]
	got := readDocs[0]
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Benchmark_StoreDocs(b *testing.B) {
	store := New()
	for i := 0; i < b.N; i++ {
		_ = store.StoreDocs(docs)
	}
}

func Benchmark_MemstoreDocs(b *testing.B) {
	store := teststore.New()
	for i := 0; i < b.N; i++ {
		_ = store.StoreDocs(docs)
	}
}
