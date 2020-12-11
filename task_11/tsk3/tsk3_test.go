// Package tsk3 реализует решение для задания 3
package tsk3

import (
	"bytes"
	"testing"
)

type Employee struct {
	name string
}

func Test_write(t *testing.T) {
	v1 := "Test"
	v2 := true
	v3 := 1
	v4 := 11.1
	v5 := Employee{name: "Test"}
	v6 := "true"

	var buf bytes.Buffer
	write(&buf, v1, v2, v3, v4, v5, v6)

	want := "Testtrue"
	got := buf.String()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
