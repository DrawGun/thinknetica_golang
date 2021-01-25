package rpcsrv

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"thinknetica_golang/task_19/pkg/structs"
)

// engine - определяет контракт движка.
type engine interface {
	Search(phrase string) ([]string, error)
}

// Service подключаемый пакет
type Service struct {
	server *RPCSrv
	port   string
}

// RPCSrv - тип данных RCP-сервера.
type RPCSrv struct {
	eng engine
}

// New позволяет создать новый объект
func New(eng engine, port string) *Service {
	s := RPCSrv{eng: eng}
	r := Service{
		server: &s,
		port:   port,
	}

	return &r
}

// Run запуск подключаемого пакета
func (s *Service) Run() {
	rpc.Register(s.server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", s.port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// Search выполняет поиск документов
func (r *RPCSrv) Search(req structs.DocumentRequest, docs *[]string) error {
	ds, err := r.eng.Search(req.SearchPhrase)
	if err != nil {
		fmt.Println(err)
		return err
	}

	*docs = ds

	return nil
}
