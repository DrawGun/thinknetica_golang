package engine

import (
	"pkg/stub"
	"testing"
)

const url = "https://yandex.ru/"
const depth = 2

func TestScan(t *testing.T) {
	s := new(stub.Scanner)

	var eng = New(url, depth)
	data, err := eng.Scan(s)
	if err != nil {
		t.Fatal(err)
	}

	got := len(data)
	want := 3

	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
