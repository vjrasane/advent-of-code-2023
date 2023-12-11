package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type pos struct {
	char int
	line int
}

func get_char_links(char rune) []pos {
	if char == 'L' {
		return []pos{{1, 0}, {0, -1}}
	}
	if char == 'F' {
		return []pos{{1, 0}, {0, 1}}
	}
	if char == 'J' {
		return []pos{{-1, 0}, {0, -1}}
	}
	if char == '7' {
		return []pos{{-1, 0}, {0, 1}}
	}
	if char == '-' {
		return []pos{{-1, 0}, {1, 0}}
	}
	if char == '|' {
		return []pos{{0, -1}, {0, 1}}
	}
	if char == 'S' {
		return []pos{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	}
	return []pos{}
}

func add(first pos, second pos) pos {
	return pos{first.char + second.char, first.line + second.line}
}

func eq(first pos, second pos) bool {
	return first.char == second.char && first.line == second.line
}

func is_connected(from pos, to pos, char rune) bool {
	for _, link := range get_char_links(char) {
		lpos := add(from, link)
		if eq(lpos, to) {
			return true
		}
	}
	return false
}

func is_in_bounds(lines []string, p pos) bool {
	if p.line < 0 || p.line >= len(lines) || p.char < 0 || p.char >= len(lines[p.line]) {
		return false
	}
	return true
}

func traverse(lines []string, current pos, previous *pos) []pos {
	fmt.Printf("%v\n", current)
	cchar := []rune(lines[current.line])[current.char]
	for _, link := range get_char_links(cchar) {
		lpos := add(current, link)
		if !is_in_bounds(lines, lpos) {
			continue
		}
		if previous != nil && eq(lpos, *previous) {
			continue
		}
		lchar := []rune(lines[lpos.line])[lpos.char]
		if is_connected(lpos, current, lchar) {
			if lchar == 'S' {
				return []pos{current}
			}
			return append([]pos{current}, traverse(lines, lpos, &current)...)
		}
	}
	return []pos{}
}

func get_start(lines []string) pos {
	for lnum, line := range lines {
		for cnum, c := range line {
			if c == 'S' {
				return pos{cnum, lnum}
			}
		}
	}
	return pos{-1, -1}
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

	start := get_start(lines)

	fmt.Println(start)

	route := traverse(lines, start, nil)

	fmt.Printf("%v %d \n", len(route), int(math.Ceil(float64(len(route))/float64(2))))

}
