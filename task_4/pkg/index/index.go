package index

import (
	"fmt"
	"pkg/crawler"
	"sort"
	"strings"
)

// Index предоставляет методы для индексирования
type Index struct {
	storage       []crawler.Document
	invertedIndex map[string][]int
}

// New создает новый экземпляр типа Index
func New(storage []crawler.Document) *Index {
	i := Index{
		storage:       storage,
		invertedIndex: map[string][]int{},
	}
	return &i
}

// Search ищет фразу по соответствию
func (i *Index) Search(phrase string) []string {
	found := []string{}
	storageLength := len(i.storage)
	ids := i.invertedIndex[phrase]

	for _, id := range ids {
		index := sort.Search(storageLength, func(ind int) bool {
			var doc = i.storage[ind]
			return doc.ID >= id
		})

		if index <= storageLength {
			document := i.storage[index]
			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
		}
	}

	return found
}

// Process заполняет индекс исходя из данных в storage
func (i *Index) Process() {
	for _, document := range i.storage {
		i.ParseAndFillIndex(document.ID, document.URL, "/", "?#:")
		i.ParseAndFillIndex(document.ID, document.Title, " ", "«()/,-")
	}
}

// ParseAndFillIndex парсит строку и заполняет invertedIndex
func (i *Index) ParseAndFillIndex(documentID int, str string, sep string, tr string) {
	lexemes := strings.Split(str, sep)

	for _, word := range lexemes {
		word = strings.Trim(word, tr)

		i.invertedIndex[word] = append(i.invertedIndex[word], documentID)
	}
}
