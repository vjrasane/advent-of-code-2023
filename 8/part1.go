package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type links struct {
	left  string
	right string
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

	inst := lines[0]

	nodes := make(map[string]links)
	for _, line := range lines[2:] {
		parts := strings.Split(line, "=")
		node := strings.TrimSpace(parts[0])
		parts = strings.Split(strings.TrimSpace(parts[1]), ", ")
		left := strings.TrimPrefix(parts[0], "(")
		right := strings.TrimSuffix(parts[1], ")")

		nodes[node] = links{left, right}
	}

	current := "AAA"

	steps := 0
	for {
		in := []rune(inst)[steps%len(inst)]
		if current == "ZZZ" {
			break
		}

		var next string
		if in == 'L' {
			next = nodes[current].left
		} else if in == 'R' {
			next = nodes[current].right
		}

		fmt.Printf("%s %c %s\n", current, in, next)

		current = next
		steps += 1
	}

	fmt.Println(steps)

}
