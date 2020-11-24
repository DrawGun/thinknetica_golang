package btree

import (
	"pkg/crawler"
	"reflect"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	btree := Tree{}

	btree.Insert(crawler.Document{ID: 1})
	btree.Insert(crawler.Document{ID: 2})
	btree.Insert(crawler.Document{ID: 3})
	btree.Insert(crawler.Document{ID: 4})
	btree.Insert(crawler.Document{ID: 5})

	var example = []int{1, 2, 3, 4, 5}
	got := reflect.DeepEqual(btree.TreeTops(), example)
	want := true
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
}

func TestTree_Search(t *testing.T) {
	btree := Tree{}

	btree.Insert(crawler.Document{ID: 1})
	btree.Insert(crawler.Document{ID: 2})
	btree.Insert(crawler.Document{ID: 3})
	btree.Insert(crawler.Document{ID: 4})
	btree.Insert(crawler.Document{ID: 5})

	got, ok := btree.Search(3)
	if !ok {
		t.Errorf("want true")
	}
	want := crawler.Document{ID: 3}

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
