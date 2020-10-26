package engine

import (
	"fmt"
	"pkg/crawler"
	"strings"
)

// Engine struct
type Engine struct {
	Crw   *crawler.Crawler
	URL   string
	depth int
}

// Scanner - интерфейс для работы с crawler
type Scanner interface {
	Scan() (map[string]string, error)
}

// New экземпляр типа Engine
func New(url string, depth int) *Engine {
	var crw = crawler.New(url, depth)
	e := Engine{
		Crw:   crw,
		URL:   url,
		depth: depth,
	}
	return &e
}

// Scan - сканирует сайт и возвращает разультат в виде ассоциативного массива
func (e *Engine) Scan(s Scanner) (map[string]string, error) {
	results, err := s.Scan()
	if err != nil {
		return make(map[string]string), err
	}

	return results, nil
}

// Search - поиск по вложению phrase в строку
func (e *Engine) Search(res map[string]string, phrase string) []string {
	var found []string

	lp := strings.ToLower(phrase)
	for k, v := range res {
		lk := strings.ToLower(k)
		lv := strings.ToLower(v)

		if strings.Contains(lk, lp) || strings.Contains(lv, lp) {
			found = append(found, fmt.Sprintf("%s -> '%s'", k, v))
		}
	}
	return found
}
