package structs

// Document - документ, веб-страница, полученная поисковым роботом.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}

// DocumentRequest - документ, веб-страница, полученная поисковым роботом.
type DocumentRequest struct {
	SearchPhrase string
}
