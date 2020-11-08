package index

import (
	"fmt"
	"index/bst"
	"pkg/crawler"
	"strings"
)

// Index предоставляет методы для индексирования
type Index struct {
	storage       bst.Tree
	invertedIndex map[string][]int
}

// New создает новый экземпляр типа Index
func New() *Index {
	i := Index{
		storage:       bst.Tree{},
		invertedIndex: map[string][]int{},
	}
	return &i
}

// Fill заполняет индекс исходя из данных в storage
func (i *Index) Fill(webDocuments *[]crawler.Document) {
	for _, document := range *webDocuments {
		i.storage.Insert(document)

		i.parseAndFill(document.ID, document.URL, "/", "?#:")
		i.parseAndFill(document.ID, document.Title, " ", "«()/,-")
	}
}

// Search ищет фразу по соответствию
func (i *Index) Search(phrase string) []string {
	found := []string{}
	ids := i.invertedIndex[phrase]

	for _, id := range ids {
		document, ok := i.storage.Search(id)

		if ok {
			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
		}
	}

	return found
}

// ParseAndFillIndex парсит строку и заполняет invertedIndex
func (i *Index) parseAndFill(documentID int, str string, sep string, tr string) {
	lexemes := strings.Split(str, sep)

	for _, word := range lexemes {
		word = strings.Trim(word, tr)

		i.invertedIndex[word] = append(i.invertedIndex[word], documentID)
	}
}
