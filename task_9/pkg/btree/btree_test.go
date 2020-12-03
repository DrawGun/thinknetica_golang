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

	var want = []int{1, 2, 3, 4, 5}
	got := btree.TreeTops()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
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
		t.Fatalf("want true")
	}
	want := crawler.Document{ID: 3}

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestElement_collect(t *testing.T) {
	type fields struct {
		left  *Element
		right *Element
		value crawler.Document
	}
	type args struct {
		tops *[]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			e.collect(tt.args.tops)
		})
	}
}
