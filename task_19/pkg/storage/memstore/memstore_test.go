package memstore

import (
	"testing"
	"thinknetica_golang/task_19/pkg/structs"
)

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

func Benchmark_StoreDocs(b *testing.B) {
	store := New()
	for i := 0; i < b.N; i++ {
		_, err := store.StoreDocs(docs)
		_ = err
	}
}
