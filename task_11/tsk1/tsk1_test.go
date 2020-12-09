// Package tsk1 реализует решение для задания 1
package tsk1

import "testing"

func TestMaxAge(t *testing.T) {
	c1 := Customer{age: 11}
	c2 := Customer{age: 22}
	c3 := Customer{age: 33}
	e1 := Employee{age: 44}
	e2 := Employee{age: 55}
	e3 := Employee{age: 66}

	want := 66
	got := MaxAge(c1, c2, c3, e1, e2, e3)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
