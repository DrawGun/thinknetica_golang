// Package btree пример реализации структуры данных "Двоичное дерево поиска"
package btree

import (
	"pkg/crawler"
	"reflect"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	bt := New()

	bt.Insert(crawler.Document{ID: 1})
	bt.Insert(crawler.Document{ID: 2})
	bt.Insert(crawler.Document{ID: 3})
	bt.Insert(crawler.Document{ID: 4})
	bt.Insert(crawler.Document{ID: 5})

	var want = []int{1, 2, 3, 4, 5}
	got := bt.Ids()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestTree_Search(t *testing.T) {
	bt := New()

	bt.Insert(crawler.Document{ID: 1})
	bt.Insert(crawler.Document{ID: 2})
	bt.Insert(crawler.Document{ID: 3})
	bt.Insert(crawler.Document{ID: 4})
	bt.Insert(crawler.Document{ID: 5})

	got, ok := bt.Search(3)
	if !ok {
		t.Fatalf("want true")
	}
	want := crawler.Document{ID: 3}

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func Benchmark_Insert(b *testing.B) {
	bt := New()
	for i := 0; i < b.N; i++ {
		bt.Insert(crawler.Document{ID: 1})
	}
}

func Benchmark_Search(b *testing.B) {
	bt := New()
	bt.Insert(crawler.Document{ID: 1})
	bt.Insert(crawler.Document{ID: 2})
	bt.Insert(crawler.Document{ID: 3})
	bt.Insert(crawler.Document{ID: 4})
	bt.Insert(crawler.Document{ID: 5})

	for i := 0; i < b.N; i++ {
		_, ok := bt.Search(3)
		_ = ok
	}
}
