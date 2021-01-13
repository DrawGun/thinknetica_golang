// Package memstore предоставляет возможность сохранить данные в памяти для тестов
package memstore

import (
	"fmt"
	"pkg/crawler"
	"sort"
	"sync"
)

// Storage - хранилище данных
type Storage struct {
	documents []crawler.Document
	id        int
	mux       *sync.Mutex
}

// New - создает новый экземпляр типа Storage
func New() *Storage {
	s := Storage{
		documents: []crawler.Document{},
		mux:       &sync.Mutex{},
	}
	return &s
}

// StoreDocs сериализует в JSON и записывает в файл массив model.Document
func (s *Storage) StoreDocs(docs []crawler.Document) ([]crawler.Document, error) {
	var updatedDocs []crawler.Document
	existedDocs := s.Docs()

	for _, doc := range docs {
		doc.ID = s.id
		updatedDocs = append(updatedDocs, doc)
		s.id++
	}

	s.documents = append(existedDocs, updatedDocs...)

	return updatedDocs, nil
}

// Docs читает из файла десериализует объекты crawler.Document
func (s *Storage) Docs() []crawler.Document {
	return s.documents
}

// Update обновляет документ по его ID
func (s *Storage) Update(doc crawler.Document) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	dCount := len(s.documents)
	index := sort.Search(dCount, func(ind int) bool {
		var d = s.documents[ind]
		return d.ID >= doc.ID
	})

	if index < dCount {
		fDoc := s.documents[index]
		if fDoc.ID == doc.ID {
			s.documents[index] = doc

			return nil
		}
	}

	return fmt.Errorf("document `%d` is not found", doc.ID)
}

// Delete удаляет документ по его ID
func (s *Storage) Delete(ID int) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	dCount := len(s.documents)
	index := sort.Search(dCount, func(ind int) bool {
		var d = s.documents[ind]
		return d.ID >= ID
	})

	if index < dCount {
		fDoc := s.documents[index]
		if fDoc.ID == ID {
			s.documents = append(s.documents[:index], s.documents[index+1:]...)

			return nil
		}
	}

	return fmt.Errorf("document `%d` is not found", ID)
}
