package index

import "thinknetica_golang/task_19/pkg/structs"

// AddSearcher - определяет контракт службы индексирования документов.
type AddSearcher interface {
	Add(docs *[]structs.Document)
	Search(string) []int
	Inspect() map[string][]int
}
