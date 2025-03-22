package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	cmd := os.Args[0]
	args := os.Args[1:]
	var num int64
	if len(args) != 1 {
		fmt.Println("Error: number required")
		fmt.Printf("Usage: %v <number>\n", filepath.Base(cmd))
		os.Exit(1)
	} else {
		var err error
		num, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Error: argument must be a number")
			os.Exit(1)
		}
	}

	results := []int64{}
	factor(num, &results)
	var display string
	for _, v := range results {
		display += fmt.Sprintf("%v ", v)
	}
	fmt.Println(display)
}
