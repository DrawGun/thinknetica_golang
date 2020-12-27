package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	cliRdr := bufio.NewReader(os.Stdin)
	connRdr := bufio.NewReader(conn)

	for {
		phrase, err := cliRdr.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if phrase == "\r\n" || phrase == "\n" {
			break
		}

		_, err = conn.Write([]byte(phrase))
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := connRdr.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(data)
	}
}
