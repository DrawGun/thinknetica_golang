// Package crawler реализует поискового робота и его интерфейс
package crawler

import "thinknetica_golang/task_19/pkg/structs"

// Scanner - интерфейс поискового робота.
type Scanner interface {
	Scan(url string, depth int) (data []structs.Document, err error)
}
