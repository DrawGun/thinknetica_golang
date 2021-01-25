// Package storage реализует хранилище и его интерфейс
package storage

import "thinknetica_golang/task_19/pkg/structs"

// Dataprocessor - определяет контракт хранилища данных.
type Dataprocessor interface {
	Docs() []structs.Document
	StoreDocs(records []structs.Document) ([]structs.Document, error)
	Update(doc structs.Document) error
	Delete(id int) error
}
