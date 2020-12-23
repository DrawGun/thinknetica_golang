package engine

import (
	"pkg/crawler/stubcrw"
	"pkg/index/hash"
	"pkg/search/btree"
	"pkg/storage/array"
	"reflect"
	"testing"
)

func TestService_PrepareData(t *testing.T) {
	store := array.New()
	ind := hash.New()
	srch := btree.New()
	eng := New(store, ind, srch)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	_, err := store.StoreDocs(docs)
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	err = eng.PrepareData()
	if err != nil {
		t.Fatalf("err = %s", err)
	}

	var want = []int{0, 1}
	got := eng.search.Ids()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_Search(t *testing.T) {
	store := array.New()
	ind := hash.New()
	srch := btree.New()
	eng := New(store, ind, srch)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	_, err := store.StoreDocs(docs)
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
