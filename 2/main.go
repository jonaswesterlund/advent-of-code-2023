package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	sum := 0
Loop:
	for i := 1; reader.Scan(); i++ {
		line := reader.Text()
		_, reveals, _ := strings.Cut(line, ": ")
		rounds := strings.Split(reveals, "; ")
		for _, v := range rounds {
			draws := strings.Split(v, ", ")
			for j := 0; j < len(draws); j++ {
				draw := strings.Split(draws[j], " ")
				number, _ := strconv.Atoi(draw[0])
				color := draw[1]
				switch {
				case color == "red" && number > 12:
					continue Loop
				case color == "green" && number > 13:
					continue Loop
				case color == "blue" && number > 14:
					continue Loop
				}
			}
		}
		sum = sum + i
	}
	fmt.Println(sum)
}
