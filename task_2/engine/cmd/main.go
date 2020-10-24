package main

import (
	"bufio"
	"fmt"
	"os"
	spider "pkg/crawler"
	"strings"
)

const url = "https://yandex.ru/"
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
		
		lt := strings.ToLower(text)
		for k, v := range titles {
			lk := strings.ToLower(k)
			lv := strings.ToLower(v)
			
			if strings.Contains(lk, lt) || strings.Contains(lv, lt) {
				fmt.Printf("%s -> '%s'\n", lk, lv)
			}
		}
	}
}
