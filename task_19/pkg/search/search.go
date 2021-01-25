// Package search реализует поиск и его интерфейс
package search

import "thinknetica_golang/task_19/pkg/structs"

// DataSearcher - определяет контракт хранилища данных.
type DataSearcher interface {
	Insert(doc structs.Document)
	Search(x int) (structs.Document, bool)
	Ids() []int
}
