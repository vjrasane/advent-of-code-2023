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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func get_next(nodes map[string]links, step int, inst string, node string) string {
	in := []rune(inst)[step%len(inst)]
	var next string
	if in == 'L' {
		next = nodes[node].left
	} else if in == 'R' {
		next = nodes[node].right
	}
	return next
}

func get_distance(nodes map[string]links, inst string, node string, nth int) []int {
	if nth <= 0 {
		return make([]int, 0)
	}

	current := node

	steps := 0
	for {
		current = get_next(nodes, steps, inst, current)
		steps += 1
		if strings.HasSuffix(current, "Z") {
			break
		}
	}

	return append([]int{steps}, get_distance(nodes, inst, current, nth-1)...)
}

func is_finished(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
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

	var current []string
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			current = append(current, node)
		}
	}

	fmt.Printf("%v\n", current)

	var distances []int
	for _, node := range current {
		distance := get_distance(nodes, inst, node, 1)

		distances = append(distances, distance[0])
		fmt.Printf("%s %v\n", node, distance)
	}

	fmt.Println(LCM(distances[0], distances[1], distances...))

}
