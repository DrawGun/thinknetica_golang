package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/index"
	"pkg/spider"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

func main() {
	var spid = spider.New()
	webData, err := Results(spid, url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

	ind := index.New(webData)
	ind.ParseStorage()

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

// Results возвращает массив просканированных ссылок
func Results(s crawler.Scanner, url string, depth int) (data []crawler.Document, err error) {
	return s.Scan(url, depth)
}
