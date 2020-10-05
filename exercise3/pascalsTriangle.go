package main

import (
	"fmt"
	"os"
	"strconv"
)

func printRow(depth int, row []int) {
	nextRow := []int{1}

	if (depth == 1) {
		fmt.Println(nextRow)
		return
	}
	
	for i := 0; i < len(row)-1; i++ {
		nextRow = append(nextRow, row[i] + row[i+1])
	}

	nextRow = append(nextRow, 1)

	if (depth == len(nextRow)) {
		fmt.Println(nextRow)
		return
	}

	printRow(depth, nextRow)
}


func main() {
	depth, _ := strconv.Atoi(os.Args[1])

	if depth > 0 {	
		printRow(depth, nil)
	}
}