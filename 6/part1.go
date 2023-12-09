package main

import (
	"bufio"
	"fmt"
	"log"
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

	times := to_nums(strings.Fields(strings.Split(lines[0], ":")[1]))
	distances := to_nums(strings.Fields(strings.Split(lines[1], ":")[1]))

	total := 1
	for racenum, time := range times {
		ways := 0
		for i := 0; i <= time; i++ {
			distance := i * (time - i)
			if distance > distances[racenum] {
				ways += 1
			}
		}
		total *= ways
	}

	fmt.Printf("%v %v %d", times, distances, total)
}
