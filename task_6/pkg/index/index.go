package index

import (
	"pkg/crawler"
	"strings"
)

// FillSearcher - интерфейс индекса
type FillSearcher interface {
	Fill(docs *[]crawler.Document) map[string][]int
}

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

// Fill заполняет индекс исходя из данных
func (i *Index) Fill(docs *[]crawler.Document) map[string][]int {
	for _, document := range *docs {
		i.fill(document.ID, document.URL, "/", "?#:")
		i.fill(document.ID, document.Title, " ", "«()/,-")
	}

	return i.InvertedIndex
}

func (i *Index) fill(documentID int, str string, sep string, tr string) {
	lexemes := strings.Split(str, sep)

	for _, word := range lexemes {
		word = strings.Trim(word, tr)

		i.InvertedIndex[word] = append(i.InvertedIndex[word], documentID)
	}
}

// Search ищет фразу по соответствию
// func (i *Index) Search(phrase string) []string {
// 	found := []string{}
// 	ids := i.InvertedIndex[phrase]

// 	for _, id := range ids {
// 		document, ok := i.storage.Search(id)

// 		if ok {
// 			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
// 		}
// 	}

// 	return found
// }
