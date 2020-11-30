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
	err1 := store.StoreDocs(docs)
	if err1 != nil {
		t.Errorf("err = %s", err1)
		return
	}

	err2 := eng.PrepareData()
	if err2 != nil {
		t.Errorf("err = %s", err2)
		return
	}

	var example = []int{0, 1}
	got := reflect.DeepEqual(eng.tree.TreeTops(), example)
	want := true
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
}

func TestService_Search(t *testing.T) {
	store := teststore.New()
	ind := hash.New()
	eng := New(store, ind)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	err1 := store.StoreDocs(docs)
	if err1 != nil {
		t.Errorf("err = %s", err1)
		return
	}

	err2 := eng.PrepareData()
	if err2 != nil {
		t.Errorf("err = %s", err2)
		return
	}

	found, err3 := eng.Search("Google")
	if err3 != nil {
		t.Errorf("err = %s", err3)
		return
	}

	var example = []string{"google.ru - Google"}
	got := reflect.DeepEqual(found, example)
	want := true
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
}
