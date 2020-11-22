package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/crawler/spider"
	"pkg/engine"
	"pkg/index/hash"
	"pkg/storage/memstore"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

// Service - сервер интернет-поисковика.
type Service struct {
	crawler crawler.Scanner
	engine  *engine.Service
}

func main() {
	service := new()

	go service.scan(url, depth)
	service.readline()
}

func new() *Service {
	store := memstore.New()
	ind := hash.New()

	s := Service{
		crawler: spider.New(),
		engine:  engine.New(store, ind),
	}

	return &s
}

// запуск процесса сканирования индексирования и сохранения информации в storage в отдельном процессе
func (srv *Service) scan(url string, depth int) {
	webDocs, err := srv.crawler.Scan(url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

	srv.engine.Storage.StoreDocs(webDocs)
	srv.engine.PrepareDataForSearch()
}

func (srv *Service) readline() {
	srv.engine.PrepareDataForSearch()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		phrase, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		phrase = strings.ReplaceAll(phrase, "\r\n", "")
		phrase = strings.ReplaceAll(phrase, "\n", "")
		if strings.Compare("", phrase) == 0 || strings.Compare("exit", phrase) == 0 {
			break
		}

		data, err := srv.engine.Search(phrase)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range data {
			fmt.Println(v)
		}
	}
}
