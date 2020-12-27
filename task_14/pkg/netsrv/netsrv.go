// Package netsrv подключаемый пакет, который обслуживает поисковые запросы
package netsrv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// engine - определяет контракт движка.
type engine interface {
	Search(phrase string) ([]string, error)
}

// Netsrv подключаемый пакет
type Netsrv struct {
	eng engine
}

// New позволяет создать новый объект
func New(eng engine) *Netsrv {
	n := Netsrv{
		eng: eng,
	}

	return &n
}

// NewListener позволяет создать новый объект с заданными настройками
func NewListener(network, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	return listener, err
}

// Run запуск подключаемого пакета
func (n *Netsrv) Run(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go n.handleConn(conn)
	}
}

func (n *Netsrv) handleConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		phrase, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		phrase = strings.ReplaceAll(phrase, "\r\n", "")
		phrase = strings.ReplaceAll(phrase, "\n", "")
		if strings.Compare("", phrase) == 0 || strings.Compare("exit", phrase) == 0 {
			break
		}

		data, err := n.eng.Search(phrase)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp = append(resp, '\n')
		_, err = conn.Write(resp)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
