// Package file предоставляет возможность сохранить данные в файл
package file

import (
	"encoding/json"
	"io/ioutil"
	"pkg/crawler"
	"time"
)

// WebDocumentsFile путь к файлу который содержит сериализованные записи crawler.Document
const WebDocumentsFile = "./../db/web_documents.json"

// IndexedDocumentsFile путь к файлу который содержит сериализованные записи crawler.Document
const IndexedDocumentsFile = "./../db/indexed_documents.json"

// Storage - служба storage.
type Storage struct {
	timestamp int64
	Test      string
}

// New создает новый экземпляр типа Engine
func New() *Storage {
	now := time.Now()
	stor := Storage{
		timestamp: now.Unix(),
	}
	return &stor
}

// Timestamp возвращает текущий записанный timestamp
func (s *Storage) Timestamp() int64 {
	return s.timestamp
}

// Write сериализует в JSON и записывает в файл массив model.Record и обратный индекс
func (s *Storage) Write(docs []crawler.Document, index map[string][]int) error {
	err := s.writeDocs(docs)
	if err != nil {
		return err
	}

	err = s.writeInvertedIndex(index)
	if err != nil {
		return err
	}

	now := time.Now()
	s.timestamp = now.Unix()

	return nil
}

// Read читает из файла десериализует объекты crawler.Document и массив index
func (s *Storage) Read() ([]crawler.Document, map[string][]int, error) {
	docs, err := s.readDocs()
	if err != nil {
		return []crawler.Document{}, map[string][]int{}, err
	}

	index, err := s.readInvertedIndex()
	if err != nil {
		return []crawler.Document{}, map[string][]int{}, err
	}

	return docs, index, nil
}

func (s *Storage) writeDocs(docs []crawler.Document) error {
	encodedDocs, err := json.Marshal(docs)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(WebDocumentsFile, encodedDocs, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) writeInvertedIndex(index map[string][]int) error {
	encodedDocs, err := json.Marshal(index)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(IndexedDocumentsFile, encodedDocs, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) readDocs() ([]crawler.Document, error) {
	docs := []crawler.Document{}

	encodedDocs, err := ioutil.ReadFile(WebDocumentsFile)
	if err != nil {
		return docs, err
	}

	err = json.Unmarshal(encodedDocs, &docs)
	if err != nil {
		return docs, err
	}

	return docs, nil
}

func (s *Storage) readInvertedIndex() (map[string][]int, error) {
	index := map[string][]int{}

	encodedIndex, err := ioutil.ReadFile(IndexedDocumentsFile)
	if err != nil {
		return index, err
	}

	err = json.Unmarshal(encodedIndex, &index)
	if err != nil {
		return index, err
	}

	return index, nil
}
