package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_index(color string) int {
	colors := []string{"red", "green", "blue"}
	for i, c := range colors {
		if c == color {
			return i
		}
	}
	return -1
}

func get_pwr(pulls []string) int {
	max_balls := []int{0, 0, 0}
	for _, pull := range pulls {
		balls := strings.Split(strings.TrimSpace(pull), ", ")
		for _, ball := range balls {
			parts := strings.Split(ball, " ")
			color := strings.TrimSpace(parts[1])
			num, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			index := get_index(color)
			if max_balls[index] < num {
				max_balls[index] = num
			}
		}
	}
	return max_balls[0] * max_balls[1] * max_balls[2]
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

		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		pulls := strings.Split(parts[1], ";")

		total += get_pwr(pulls)

		fmt.Printf("%d\n", id)
		// total += value
	}
	fmt.Println(total)
}
