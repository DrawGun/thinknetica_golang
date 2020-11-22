package memstore

import (
	"encoding/json"
	"io/ioutil"
	"pkg/crawler"
)

// WebDocumentsFile путь к файлу который содержит сериализованные записи crawler.Document
const WebDocumentsFile = "./../../db/web_documents.json"

// DB - хранилище данных
type DB struct{}

// New - создает новый экземпляр типа DB
func New() *DB {
	db := DB{}
	return &db
}

// StoreDocs сериализует в JSON и записывает в файл массив model.Document
func (db *DB) StoreDocs(docs []crawler.Document) error {
	err := db.writeDocs(docs)
	if err != nil {
		return err
	}

	return nil
}

// Docs читает из файла десериализует объекты crawler.Document
func (db *DB) Docs() ([]crawler.Document, error) {
	docs, err := db.readDocs()
	if err != nil {
		return []crawler.Document{}, err
	}

	return docs, nil
}

func (db *DB) writeDocs(docs []crawler.Document) error {
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

func (db *DB) readDocs() ([]crawler.Document, error) {
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
