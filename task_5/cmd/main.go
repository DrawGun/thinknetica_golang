package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/index"
	"pkg/spider"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

func main() {
	var spid = spider.New()
	webDocuments, err := spid.Scan(url, depth)

	if err != nil {
		fmt.Println(err)
		return
	}

	ind := index.New()
	ind.Fill(&webDocuments)

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

		data := ind.Search(phrase)

		for _, v := range data {
			fmt.Println(v)
		}
	}
}
