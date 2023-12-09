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

	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))

	ways := 0
	for i := 0; i <= time; i++ {
		d := i * (time - i)
		if d > distance {
			ways += 1
		}
	}

	fmt.Printf("%d %d %d", time, distance, ways)
}
