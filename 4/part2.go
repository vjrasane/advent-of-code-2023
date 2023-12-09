package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func to_nums(values []string) []int {
	var nums []int
	for _, value := range values {
		num, _ := strconv.Atoi(strings.TrimSpace(value))

		nums = append(nums, num)
	}
	return nums
}

func is_winning(num int, winning []int) bool {
	for _, win := range winning {
		if num == win {
			return true
		}

	}
	return false
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	slices.Reverse(lines)
	var counts []int
	for linenum, line := range lines {
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], " | ")
		winning := to_nums(strings.Fields(parts[0]))
		numbers := to_nums(strings.Fields(parts[1]))

		count := 0
		for _, num := range numbers {
			if is_winning(num, winning) {
				count += 1
			}
		}

		combined := 1
		for i := 1; i <= count; i++ {
			combined += counts[linenum-i]
		}

		counts = append(counts, combined)
	}

	total := 0
	for _, count := range counts {
		total += count
	}

	fmt.Println(total)
}
