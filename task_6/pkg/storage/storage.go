// Package storage реализует хранилище и его интерфейс
package storage

import "pkg/crawler"

// Dataprocessor - интерфейс хранилища
type Dataprocessor interface {
	Write(records []crawler.Document, index map[string][]int) error
	Read() ([]crawler.Document, map[string][]int, error)
	Timestamp() int64
}
