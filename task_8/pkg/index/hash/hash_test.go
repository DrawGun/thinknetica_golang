package hash

import (
	"pkg/crawler/stubcrw"
	"testing"
)

func TestIndex_Add(t *testing.T) {
	scanner := stubcrw.New()
	data, _ := scanner.Scan("yandex.ru", 1)

	ind := New()
	ind.Add(&data)

	got := len(ind.InvertedIndex)
	want := 4

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestIndex_Search(t *testing.T) {
	var data = map[string][]int{
		"Google":    []int{1},
		"google.ru": []int{1},
		"yandex.ru": []int{0},
		"Яндекс":    []int{0},
	}
	ind := New()
	ind.InvertedIndex = data

	got := ind.Search("Google")
	want := []int{1}

	if got[0] != want[0] {
		t.Errorf("got %d; want %d", got, want)
	}
}
