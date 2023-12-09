package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		// card := parts[0]
		parts = strings.Split(parts[1], " | ")
		winning := to_nums(strings.Fields(parts[0]))
		numbers := to_nums(strings.Fields(parts[1]))

		fmt.Printf("%v | [", winning)

		count := 0
		for _, num := range numbers {
			if is_winning(num, winning) {
				fmt.Printf("<%d>", num)
				count += 1
			} else {

				fmt.Printf("%d ", num)
			}
		}
		points := 0
		if count > 0 {
			points = int(math.Pow(2, float64(count-1)))
		}
		total += points

		fmt.Printf("] | %d %d\n", count, points)
		// fmt.Printf("%v %v %d %d\n", winning, numbers, count, points)

	}
	fmt.Println(total)
}
