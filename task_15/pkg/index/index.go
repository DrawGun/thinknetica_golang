package index

import "pkg/crawler"

// AddSearcher - определяет контракт службы индексирования документов.
type AddSearcher interface {
	Add(docs *[]crawler.Document)
	Search(string) []int
	Aggregate() map[string][]int
}
