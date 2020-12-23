package memstore

import (
	"encoding/json"
	"io/ioutil"
	"pkg/crawler"
)

// webDocumentsFile путь к файлу который содержит сериализованные записи crawler.Document
var webDocumentsFile = "./web_documents.json"

// DB - хранилище данных
type DB struct {
	id int
}

// New - создает новый экземпляр типа DB
func New() *DB {
	db := DB{}
	return &db
}

// StoreDocs сериализует в JSON и записывает в файл массив model.Document
func (db *DB) StoreDocs(docs []crawler.Document) ([]crawler.Document, error) {
	updatedDocs, err := db.writeDocs(docs)
	if err != nil {
		return []crawler.Document{}, err
	}

	return updatedDocs, nil
}

// Docs читает из файла десериализует объекты crawler.Document
func (db *DB) Docs() []crawler.Document {
	docs := db.readDocs()

	return docs
}

func (db *DB) writeDocs(docs []crawler.Document) ([]crawler.Document, error) {
	var updatedDocs []crawler.Document
	existedDocs := db.Docs()

	for _, doc := range docs {
		doc.ID = db.id
		updatedDocs = append(updatedDocs, doc)
		db.id++
	}

	commonDocs := append(existedDocs, updatedDocs...)
	encodedDocs, err := json.Marshal(commonDocs)
	if err != nil {
		return []crawler.Document{}, err
	}

	err = ioutil.WriteFile(webDocumentsFile, encodedDocs, 0666)
	if err != nil {
		return []crawler.Document{}, err
	}

	return updatedDocs, nil
}

func (db *DB) readDocs() []crawler.Document {
	docs := []crawler.Document{}

	encodedDocs, err := ioutil.ReadFile(webDocumentsFile)
	if err != nil {
		return docs
	}

	err = json.Unmarshal(encodedDocs, &docs)
	if err != nil {
		return docs
	}

	return docs
}
