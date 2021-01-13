package main

import (
	"fmt"
	"pkg/api"
	"pkg/crawler"
	"pkg/crawler/spider"
	"pkg/engine"
	"pkg/index/hash"
	"pkg/search/btree"
	"pkg/storage/memstore"
	"pkg/webapp"
	"sync"
)

const scanerWorkers = 10

// Service - сервер интернет-поисковика.
type Service struct {
	api     *api.Service
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
	service.api.Run()

	wa := webapp.New(service.engine.Index, service.engine.Storage, ":8000")
	wa.Run()
}

func new() *Service {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()
	eng := engine.New(store, ind, srch)
	api := api.New(eng, ":3000")

	s := Service{
		crawler: spider.New(),
		engine:  eng,
		sites:   sites,
		api:     api,
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

			srv.engine.UpdateDocuments(updatedDocs)
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
