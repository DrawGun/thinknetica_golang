package crawler

// Scanner - интерфейс поискового робота.
type Scanner interface {
	Scan() ([]Document, error)
}

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}
