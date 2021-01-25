package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
	"thinknetica_golang/task_19/pkg/structs"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text = strings.ReplaceAll(text, "\r\n", "")
		text = strings.ReplaceAll(text, "\n", "")

		if strings.Compare("", text) == 0 || strings.Compare("exit", text) == 0 {
			break
		}

		docs, err := search(text)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, d := range docs {
			fmt.Println(d)
		}
	}
}

func search(text string) ([]string, error) {
	client, err := rpc.DialHTTP("tcp", ":3001")
	if err != nil {
		return []string{}, err
	}
	defer client.Close()
	req := &structs.DocumentRequest{SearchPhrase: text}
	var docs []string
	err = client.Call("RPCSrv.Search", req, &docs)
	if err != nil {
		return []string{}, err
	}

	return docs, nil
}
