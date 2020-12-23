package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/crawler/spider"
	"pkg/engine"
	"pkg/index/hash"
	"pkg/search/btree"
	"pkg/storage/memstore"
	"strings"
	"sync"
)

const scanerWorkers = 10

// Service - сервер интернет-поисковика.
type Service struct {
	crawler crawler.Scanner
	engine  *engine.Service
	sites   []siteToScan
}

type siteToScan struct {
	url   string
	depth int
}

var sites = []siteToScan{
	{"https://habr.com", 1},
	{"https://go.dev", 1},
	{"https://golang.org/", 1},
	{"https://learn.go.dev/", 1},
	{"https://www.ruby-lang.org/", 1},
	{"http://rusrails.ru/", 1},
	{"https://www.postgresql.org/", 1},
	{"https://redis.io/", 1},
	{"https://memcached.org/", 1},
	{"https://www.rabbitmq.com/", 1},
	{"https://kafka.apache.org/", 1},
	{"https://prometheus.io/", 1},
}

func main() {
	service := new()

	go service.scan()
	service.readline()
}

func new() *Service {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()

	s := Service{
		crawler: spider.New(),
		engine:  engine.New(store, ind, srch),
		sites:   sites,
	}

	return &s
}

func (srv *Service) scan() {
	scannerJobs := make(chan siteToScan)
	rawDocs := make(chan []crawler.Document)
	wg := sync.WaitGroup{}
	wg.Add(scanerWorkers)

	for i := 0; i < scanerWorkers; i++ {
		go srv.scanWorker(&wg, scannerJobs, rawDocs)
	}

	go func(ch chan<- siteToScan) {
		for _, site := range srv.sites {
			ch <- site
		}
		close(ch)
	}(scannerJobs)

	go func(ch <-chan []crawler.Document) {
		for docs := range ch {
			updatedDocs, err := srv.engine.Storage.StoreDocs(docs)
			if err != nil {
				fmt.Println(err)
				continue
			}

			srv.engine.UpdateData(updatedDocs)
		}
	}(rawDocs)

	wg.Wait()
	close(rawDocs)
}

func (srv *Service) scanWorker(wg *sync.WaitGroup, jobs <-chan siteToScan, results chan<- []crawler.Document) {
	defer wg.Done()

	for job := range jobs {
		webDocs, err := srv.crawler.Scan(job.url, job.depth)

		if err != nil {
			fmt.Println(err)
			continue
		}
		results <- webDocs
	}
}

func (srv *Service) readline() {
	err := srv.engine.PrepareData()

	if err != nil {
		fmt.Println(err)
		return
	}

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
