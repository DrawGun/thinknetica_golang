package main

import (
	"bufio"
	"fmt"
	"os"
	spider "pkg/crawler"
	"strings"
)

const url = "https://pikabu.ru/"
const depth = 2

func main() {
	fmt.Printf("Scanning '%s'...\n", url)

	titles, err := spider.Scan(url, depth)

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")

		if strings.Compare("", text) == 0 || strings.Compare("exit", text) == 0 {
			break
		}

		for k, v := range titles {
			if strings.Contains(k, text) || strings.Contains(v, text) {
				fmt.Printf("%s -> '%s'\n", k, v)
			}
		}
	}
}
