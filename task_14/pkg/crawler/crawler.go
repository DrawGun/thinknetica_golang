// Package crawler реализует поискового робота и его интерфейс
package crawler

// Scanner - интерфейс поискового робота.
type Scanner interface {
	Scan(url string, depth int) (data []Document, err error)
}

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}
