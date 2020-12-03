package engine

import (
	"pkg/crawler/stubcrw"
	"pkg/index/hash"
	"pkg/storage/teststore"
	"reflect"
	"testing"
)

func TestService_PrepareData(t *testing.T) {
	store := teststore.New()
	ind := hash.New()
	eng := New(store, ind)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	err := store.StoreDocs(docs)
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	err = eng.PrepareData()
	if err != nil {
		t.Fatalf("err = %s", err)
	}

	var want = []int{0, 1}
	got := eng.tree.TreeTops()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_Search(t *testing.T) {
	store := teststore.New()
	ind := hash.New()
	eng := New(store, ind)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	err := store.StoreDocs(docs)
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	err = eng.PrepareData()
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	got, err := eng.Search("Google")
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	var want = []string{"google.ru - Google"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}
