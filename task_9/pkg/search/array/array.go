// Package array для тестирования производительности
package array

import (
	"pkg/crawler"
	"sort"
)

// Service предоставляет методы для поиска данных
type Service struct {
	arr []crawler.Document
}

// New создает новый экземпляр типа Tree
func New() *Service {
	s := Service{}
	return &s
}

// Insert - вставка элемента в дерево
func (s *Service) Insert(doc crawler.Document) {
	s.arr = append(s.arr, doc)
}

// Search - поиск значения в дереве, выдаёт true если найдено, иначе false
func (s *Service) Search(x int) (crawler.Document, bool) {
	arrLength := len(s.arr)
	index := sort.Search(arrLength, func(ind int) bool {
		var doc = s.arr[ind]
		return doc.ID >= x
	})

	if index <= arrLength {
		document := s.arr[index]
		return document, true
	}

	return crawler.Document{}, false
}

// Ids - возвращает массив вершин дерева
func (s *Service) Ids() []int {
	ids := []int{}
	for _, document := range s.arr {
		ids = append(ids, document.ID)
	}

	return ids
}
