// Package teststore предоставляет возможность сохранить данные в памяти для тестов
package teststore

import "pkg/crawler"

// Storage - хранилище данных
type Storage struct {
	documents []crawler.Document
}

// New - создает новый экземпляр типа Storage
func New() *Storage {
	s := Storage{
		documents: []crawler.Document{},
	}
	return &s
}

// StoreDocs сериализует в JSON и записывает в файл массив model.Document
func (s *Storage) StoreDocs(docs []crawler.Document) error {
	s.documents = docs

	return nil
}

// Docs читает из файла десериализует объекты crawler.Document
func (s *Storage) Docs() []crawler.Document {
	return s.documents
}
