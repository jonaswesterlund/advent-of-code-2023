package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	sum := 0
	for reader.Scan() {
		line := reader.Text()
		numbers := make([]int, 0)
		for _, char := range line {
			if num, err := strconv.Atoi(string(char)); err == nil {
				numbers = append(numbers, num)

			}
		}
		value, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1]))
		sum = sum + value
	}
	fmt.Print(sum)
}
