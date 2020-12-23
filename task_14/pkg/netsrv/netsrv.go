// Package netsrv подключаемый пакет, который обслуживает поисковые запросы
package netsrv

import (
	"bufio"
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
	eng  engine
	netw string
	addr string
}

// New позволяет создать новый объект с заданными настройками
func New(eng engine, network, address string) *Netsrv {
	n := Netsrv{
		eng:  eng,
		netw: network,
		addr: address,
	}

	return &n
}

// Run запуск подключаемого пакета
func (n *Netsrv) Run() {
	listener, err := net.Listen(n.netw, n.addr)
	if err != nil {
		fmt.Println(err)
		return
	}

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

		for _, v := range data {
			msg := []byte(v + "\n")

			_, err = conn.Write(msg)
			if err != nil {
				return
			}
		}
	}
}
