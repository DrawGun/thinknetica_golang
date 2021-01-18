// Package storage реализует хранилище и его интерфейс
package storage

import "pkg/crawler"

// Dataprocessor - определяет контракт хранилища данных.
type Dataprocessor interface {
	Docs() []crawler.Document
	StoreDocs(records []crawler.Document) ([]crawler.Document, error)
	Update(doc crawler.Document) error
	Delete(id int) error
}
