package main

import (
	"flag"
	"fmt"
	"pkg/fibonacci"
)

func main() {
	var nFlag = flag.Int("n", 0, "Something about flag n")
	flag.Parse()

	var n = *nFlag
	if n < 0 || n > 20 {
		fmt.Println("Invalid value of flag -n. It can be chosen in the range 0..20")
		return
	}

	result := fibonacci.Calculate(n)
	fmt.Println(result)
}
