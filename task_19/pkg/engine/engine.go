package engine

import (
	"fmt"
	"thinknetica_golang/task_19/pkg/index"
	"thinknetica_golang/task_19/pkg/search"
	"thinknetica_golang/task_19/pkg/storage"
	"thinknetica_golang/task_19/pkg/structs"
)

// Service - поисковый движок.
type Service struct {
	Storage storage.Dataprocessor
	Index   index.AddSearcher
	search  search.DataSearcher
}

// New создает новый экземпляр типа Service
func New(store storage.Dataprocessor, ind index.AddSearcher, srch search.DataSearcher) *Service {
	srv := Service{
		Storage: store,
		Index:   ind,
		search:  srch,
	}
	return &srv
}

// UpdateDocuments подготавливает данные для поиска
func (srv *Service) UpdateDocuments(docs []structs.Document) {
	srv.Index.Add(&docs)

	for _, document := range docs {
		srv.search.Insert(document)
	}
}

// Search осуществляет поиск по слову
func (srv *Service) Search(phrase string) ([]string, error) {
	found := []string{}

	ids := srv.Index.Search(phrase)

	for _, id := range ids {
		document, ok := srv.search.Search(id)

		if ok {
			found = append(found, fmt.Sprintf("%s - %s", document.URL, document.Title))
		}
	}

	return found, nil
}

// DocumentByID осуществляет поиск документа по id
func (srv *Service) DocumentByID(id int) (structs.Document, bool) {
	return srv.search.Search(id)
}

// UpdateDocument обновляет документ
func (srv *Service) UpdateDocument(doc structs.Document) error {
	err := srv.Storage.Update(doc)
	return err
}

// DeleteDocument удаляет документ
func (srv *Service) DeleteDocument(id int) error {
	err := srv.Storage.Delete(id)
	return err
}
