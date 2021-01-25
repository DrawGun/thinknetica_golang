package engine

import (
	"reflect"
	"testing"
	"thinknetica_golang/task_19/pkg/crawler/stubcrw"
	"thinknetica_golang/task_19/pkg/index/hash"
	"thinknetica_golang/task_19/pkg/search/btree"
	"thinknetica_golang/task_19/pkg/storage/memstore"
)

func TestService_PrepareData(t *testing.T) {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()
	eng := New(store, ind, srch)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	_, err := store.StoreDocs(docs)
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	eng.UpdateDocuments(docs)

	var want = []int{0, 1}
	got := eng.search.Ids()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestService_Search(t *testing.T) {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()
	eng := New(store, ind, srch)

	scanner := stubcrw.New()
	docs, _ := scanner.Scan("yandex.ru", 1)
	_, err := store.StoreDocs(docs)
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	eng.UpdateDocuments(docs)

	got, err := eng.Search("Google")
	if err != nil {
		t.Fatalf("err = %q", err)
	}

	var want = []string{"google.ru - Google"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}
