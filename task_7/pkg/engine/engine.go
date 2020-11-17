package engine

import (
	"fmt"
	"pkg/btree"
	"pkg/index"
	"pkg/storage"
)

// Service - поисковый движок.
type Service struct {
	Storage storage.Dataprocessor
	Index   index.AddSearcher
	tree    btree.Tree
}

// New создает новый экземпляр типа Service
func New(store storage.Dataprocessor, ind index.AddSearcher) *Service {
	srv := Service{
		Storage: store,
		Index:   ind,
		tree:    btree.Tree{},
	}
	return &srv
}

// PrepareDataForSearch подготавливает данные для поиска
func (srv *Service) PrepareDataForSearch() error {
	docs, err := srv.Storage.Docs()
	if err != nil {
		return err
	}

	srv.Index.Add(&docs)

	for _, document := range docs {
		srv.tree.Insert(document)
	}

	return nil
}

// Search осуществляет поиск по слову
func (srv *Service) Search(phrase string) ([]string, error) {
	found := []string{}

	ids := srv.Index.Search(phrase)

	for _, id := range ids {
		document, ok := srv.tree.Search(id)

		if ok {
			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
		}
	}

	return found, nil
}
