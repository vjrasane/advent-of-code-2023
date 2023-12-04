package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_max() map[string]int {
	balls := make(map[string]int)
	balls["blue"] = 14
	balls["red"] = 12
	balls["green"] = 13
	return balls
}

func is_possible(pulls []string) bool {
	max_balls := get_max()
	for _, pull := range pulls {
		balls := strings.Split(strings.TrimSpace(pull), ", ")
		for _, ball := range balls {
			parts := strings.Split(ball, " ")
			color := strings.TrimSpace(parts[1])
			num, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			if max_balls[color] < num {
				return false
			}
		}
	}
	return true
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

		if is_possible(pulls) {
			total += id
		}

		fmt.Printf("%d\n", id)
		// total += value
	}
	fmt.Println(total)
}
