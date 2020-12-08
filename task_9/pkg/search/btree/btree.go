// Package btree пример реализации структуры данных "Двоичное дерево поиска"
package btree

import (
	"fmt"
	"pkg/crawler"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	root *Element
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	value       crawler.Document
}

// New создает новый экземпляр типа Tree
func New() *Tree {
	t := Tree{}
	return &t
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(doc crawler.Document) {
	newEl := &Element{value: doc}

	if t.root == nil {
		t.root = newEl
		return
	}
	insert(t.root, newEl)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.value.ID < node.value.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.value.ID >= node.value.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search - поиск значения в дереве, выдаёт true если найдено, иначе false
func (t *Tree) Search(x int) (crawler.Document, bool) {
	return search(t.root, x)
}
func search(el *Element, x int) (crawler.Document, bool) {
	if el == nil {
		return crawler.Document{}, false
	}
	if el.value.ID == x {
		return el.value, true
	}
	if el.value.ID < x {
		return search(el.right, x)
	}
	return search(el.left, x)
}

// String - реализуем интерфейс Stringer для функций печати пакета fmt
func (t Tree) String() string {
	return prettyPrint(t.root, 0)
}

// prettyPrint печатает дерево в виде дерева :)
func prettyPrint(e *Element, spaces int) (res string) {
	if e == nil {
		return res
	}

	spaces++
	res += prettyPrint(e.right, spaces)
	for i := 0; i < spaces; i++ {
		res += fmt.Sprint("\t")
	}
	res += fmt.Sprintf("%d\n", e.value.ID)
	res += prettyPrint(e.left, spaces)

	return res
}

// Ids - возвращает массив вершин дерева
func (t Tree) Ids() []int {
	ids := []int{}
	t.root.collect(&ids)

	return ids
}

func (e *Element) collect(ids *[]int) {
	if e == nil {
		return
	}

	*ids = append(*ids, e.value.ID)

	e.left.collect(ids)
	e.right.collect(ids)
}
