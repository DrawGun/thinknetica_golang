package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/engine"
	"strings"
)

const url = "https://yandex.ru/"
const depth = 2

func main() {
	var eng = engine.New(url, depth)

	results, err := eng.Results()
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

		data := eng.Search(results, phrase)

		for _, v := range data {
			fmt.Println(v)
		}
	}
}
