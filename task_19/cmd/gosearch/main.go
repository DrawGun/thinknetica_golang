package main

import (
	"fmt"
	"thinknetica_golang/task_19/pkg/api"
	"thinknetica_golang/task_19/pkg/crawler"
	"thinknetica_golang/task_19/pkg/crawler/spider"
	"thinknetica_golang/task_19/pkg/engine"
	"thinknetica_golang/task_19/pkg/index/hash"
	"thinknetica_golang/task_19/pkg/rpcsrv"
	"thinknetica_golang/task_19/pkg/search/btree"
	"thinknetica_golang/task_19/pkg/storage/memstore"
	"thinknetica_golang/task_19/pkg/structs"
	"thinknetica_golang/task_19/pkg/webapp"

	"sync"
)

const scanerWorkers = 10

// Service - сервер интернет-поисковика.
type Service struct {
	api     *api.Service
	crawler crawler.Scanner
	engine  *engine.Service
	sites   []siteToScan
	rpcsrv  *rpcsrv.Service
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
	go service.api.Run()
	go service.rpcsrv.Run()

	wa := webapp.New(service.engine.Index, service.engine.Storage, ":8000")
	wa.Run()
}

func new() *Service {
	store := memstore.New()
	ind := hash.New()
	srch := btree.New()
	eng := engine.New(store, ind, srch)
	api := api.New(eng, ":3000")
	rpcsrv := rpcsrv.New(eng, ":3001")

	s := Service{
		crawler: spider.New(),
		engine:  eng,
		sites:   sites,
		api:     api,
		rpcsrv:  rpcsrv,
	}

	return &s
}

func (srv *Service) scan() {
	scannerJobs := make(chan siteToScan)
	rawDocs := make(chan []structs.Document)
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

	go func(ch <-chan []structs.Document) {
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

func (srv *Service) scanWorker(wg *sync.WaitGroup, jobs <-chan siteToScan, results chan<- []structs.Document) {
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
