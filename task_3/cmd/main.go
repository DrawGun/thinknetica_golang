package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/spider"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

func main() {
	var spid = spider.New()

	results, err := Results(spid, url, depth)
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

		data := search(results, phrase)

		for _, v := range data {
			fmt.Println(v)
		}
	}
}

// Results возвращает массив по вхождению строки
func Results(s crawler.Scanner, url string, depth int) (data []crawler.Document, err error) {
	return s.Scan(url, depth)
}

func search(res []crawler.Document, phrase string) []string {
	var found []string

	lp := strings.ToLower(phrase)
	for _, doc := range res {
		lt := strings.ToLower(doc.Title)
		lu := strings.ToLower(doc.URL)

		if strings.Contains(lt, lp) || strings.Contains(lu, lp) {
			found = append(found, fmt.Sprintf("%s -> '%s'", doc.Title, doc.URL))
		}
	}
	return found
}
