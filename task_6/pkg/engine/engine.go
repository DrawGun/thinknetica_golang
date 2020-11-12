// Package engine - по задумке основной пакет, разруливающий все взаимодействие в системе
package engine

import (
	"fmt"
	"pkg/btree"
	"pkg/crawler"
	"pkg/index"
	"pkg/storage/file"
)

// Engine предоставляет основные методы
type Engine struct {
	Crawler crawler.Scanner
	Storage *file.Storage
	Index   *index.Index
	tree    btree.Tree
}

// New создает новый экземпляр типа Engine
func New(crw crawler.Scanner, stor *file.Storage, ind *index.Index) *Engine {
	eng := Engine{
		Crawler: crw,
		Storage: stor,
		Index:   ind,
		tree:    btree.Tree{},
	}
	return &eng
}

// Search осуществляет поиск по слову
func (e *Engine) Search(phrase string) ([]string, error) {
	found := []string{}

	if e.tree.Timestamp != e.Storage.Timestamp() {
		e.prepareDataForSearch()
	}

	ids := e.Index.InvertedIndex[phrase]
	for _, id := range ids {
		document, ok := e.tree.Search(id)

		if ok {
			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
		}
	}

	return found, nil
}

func (e *Engine) prepareDataForSearch() error {
	docs, index, err := e.Storage.Read()
	if err != nil {
		return err
	}

	e.Index.InvertedIndex = index

	for _, document := range docs {
		e.tree.Insert(document)
	}

	e.tree.Timestamp = e.Storage.Timestamp()

	return nil
}
