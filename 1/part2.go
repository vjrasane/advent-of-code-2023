package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_words() map[string]int {
	words := make(map[string]int)
	words["zero"] = 0
	words["one"] = 1
	words["two"] = 2
	words["three"] = 3
	words["four"] = 4
	words["five"] = 5
	words["six"] = 6
	words["seven"] = 7
	words["eight"] = 8
	words["nine"] = 9
	words["ten"] = 10
	return words
}

func get_nums(line string) []int {
	words := get_words()
	var nums []int
	for i, char := range line {
		value, err := strconv.Atoi(string(char))
		if err == nil {
			nums = append(nums, value)
		}

		remaining := strings.TrimSpace(line[i:len(line)])
		for word, digit := range words {
			if strings.HasPrefix(remaining, word) {
				nums = append(nums, digit)
			}
		}
	}
	return nums
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func get_line_num(line string) int {
	fmt.Println(line)
	nums := get_nums(line)

	combined, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
	if err != nil {
		log.Fatal(err)
	}
	return combined
}

func get_result() int {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		total += get_line_num(scanner.Text())
	}
	return total
}

func main() {
	result := get_result()
	fmt.Println(result)
}
