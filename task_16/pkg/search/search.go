// Package search реализует поиск и его интерфейс
package search

import "pkg/crawler"

// DataSearcher - определяет контракт хранилища данных.
type DataSearcher interface {
	Insert(doc crawler.Document)
	Search(x int) (crawler.Document, bool)
	Ids() []int
}
