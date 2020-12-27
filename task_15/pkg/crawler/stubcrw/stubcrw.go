// Package stubcrw реализует заглушку сканера.
package stubcrw

import "pkg/crawler"

// StubCrw имитирует Crawler
type StubCrw struct{}

// New создает новый экземпляр
func New() *StubCrw {
	stub := StubCrw{}
	return &stub
}

// Scan возвращает заранее подготовленный набор данных
func (s *StubCrw) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			ID:    0,
			URL:   "yandex.ru",
			Title: "Яндекс",
		},
		{
			ID:    1,
			URL:   "google.ru",
			Title: "Google",
		},
	}

	return data, nil
}
