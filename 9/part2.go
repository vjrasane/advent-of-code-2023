package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func is_zeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func extrapolate(history []int) int {
	if is_zeroes(history) {
		return 0
	}

	var changes []int
	prev := history[0]
	for _, num := range history[1:] {
		changes = append(changes, num-prev)
		prev = num
	}

	change := extrapolate(changes)
	value := history[0] - change
	// fmt.Printf("%d %v %d %v\n", value, history, change, changes)
	return value
}

func to_nums(values []string) []int {
	var nums []int
	for _, value := range values {
		num, _ := strconv.Atoi(strings.TrimSpace(value))

		nums = append(nums, num)
	}
	return nums
}
func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var histories [][]int
	for scanner.Scan() {
		line := scanner.Text()

		history := to_nums(strings.Fields(line))
		histories = append(histories, history)
	}

	total := 0
	for _, history := range histories {
		value := extrapolate(history)

		// fmt.Printf("%v %d\n", history, value)
		total += value
	}

	fmt.Println(total)

}
