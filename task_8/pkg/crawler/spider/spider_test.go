package spider

import (
	"testing"
)

func TestService_Scan(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	const url = "https://yandex.ru"
	const depth = 1
	spid := New()
	data, err := spid.Scan(url, depth)
	if err != nil {
		t.Errorf("err = %s", err)
		return
	}

	got := len(data)
	want := 1
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
