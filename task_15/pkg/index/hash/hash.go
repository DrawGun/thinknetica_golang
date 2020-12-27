package hash

import (
	"pkg/crawler"
	"strings"
)

// Index предоставляет методы для индексирования
type Index struct {
	InvertedIndex map[string][]int
}

// New создает новый экземпляр типа Index
func New() *Index {
	i := Index{
		InvertedIndex: map[string][]int{},
	}
	return &i
}

// Add заполняет индекс исходя из данных
func (i *Index) Add(docs *[]crawler.Document) {
	for _, document := range *docs {
		i.fill(document.ID, document.URL, "/", "?#:")
		i.fill(document.ID, document.Title, " ", "«()/,-")
	}
}

// Search возвращает id индексов по совпадению фразы
func (i *Index) Search(phrase string) []int {
	return i.InvertedIndex[phrase]
}

func (i *Index) fill(documentID int, str string, sep string, tr string) {
	lexemes := strings.Split(str, sep)

	for _, word := range lexemes {
		word = strings.Trim(word, tr)

		i.InvertedIndex[word] = append(i.InvertedIndex[word], documentID)
	}
}

// Aggregate возвращает полностью InvertedIndex
func (i *Index) Aggregate() map[string][]int {
	return i.InvertedIndex
}
