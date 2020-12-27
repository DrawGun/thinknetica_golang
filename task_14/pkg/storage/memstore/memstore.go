// Package memstore предоставляет возможность сохранить данные в памяти для тестов
package memstore

import "pkg/crawler"

// Storage - хранилище данных
type Storage struct {
	documents []crawler.Document
	id        int
}

// New - создает новый экземпляр типа Storage
func New() *Storage {
	s := Storage{
		documents: []crawler.Document{},
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
