package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler/spider"
	"pkg/engine"
	"pkg/index"
	"pkg/storage/file"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

func main() {
	spid := spider.New()
	file := file.New()
	ind := index.New()

	eng := engine.New(spid, file, ind)

	go backgroundUpdates(eng)

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

		data, err := eng.Search(phrase)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range data {
			fmt.Println(v)
		}
	}
}

// запуск процесса сканирования индексирования и сохранения информации в storage в отдельном процессе
func backgroundUpdates(eng *engine.Engine) {
	webDocs, err := eng.Crawler.Scan(url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

	invInd := eng.Index.Fill(&webDocs)
	eng.Storage.Write(webDocs, invInd)
}
