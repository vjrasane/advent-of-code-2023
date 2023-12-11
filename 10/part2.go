package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func get_char_rotation(char rune, p pos) pos {
	if char == '7' || char == 'L' {
		return pos{-p.line, -p.char}
	}
	if char == 'F' || char == 'J' {
		return pos{p.line, p.char}
	}

	return p
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
	// fmt.Printf("%v\n", current)
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
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
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

func print_lines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func draw_at(lines []string, p pos, char rune) {
	lines[p.line] = replaceAtIndex(lines[p.line], char, p.char)
}

func draw(lines []string, d func(pos, rune) rune) []string {
	var drawn []string
	for lnum, line := range lines {
		str := ""
		for cnum, char := range line {
			current := pos{cnum, lnum}
			str += string(d(current, char))
		}
		drawn = append(drawn, str)
	}
	return drawn
}

func get_next(lines []string, current pos, prev pos) pos {
	cchar := char_at(lines, current)
	for _, link := range get_char_links(cchar) {
		lpos := add(link, current)
		if eq(lpos, prev) {
			continue
		}
		return lpos
	}
	return pos{-1, -1}
}

func char_at(lines []string, p pos) rune {
	return []rune(lines[p.line])[p.char]
}

func mark_inside(lines []string, current pos, visited [][]bool) {
	if !is_in_bounds(lines, current) {
		return
	}
	if visited[current.line][current.char] {
		return
	}

	if char_at(lines, current) != '.' {
		return
	}
	draw_at(lines, current, 'I')
	visited[current.line][current.char] = true

	for _, link := range append(get_char_links('S'), []pos{
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}...) {
		lpos := add(link, current)
		mark_inside(lines, lpos, visited)
	}
}

func mark(lines []string, current pos, inside pos, prev pos, visited [][]bool) {
	next := get_next(lines, current, prev)
	if visited[next.line][next.char] {
		return
	}

	ipos := add(next, inside)
	mark_inside(lines, ipos, visited)

	if char_at(lines, next) == 'S' {
		return
	}

	nchar := char_at(lines, next)
	rotated := get_char_rotation(nchar, inside)

	if !eq(inside, rotated) {
		ipos = add(next, rotated)

		mark_inside(lines, ipos, visited)
	}

	visited[next.line][next.char] = true
	mark(lines, next, rotated, current, visited)
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

	var visited [][]bool
	for _, line := range lines {
		vline := make([]bool, len(line))
		for i := range line {
			vline[i] = false
		}
		visited = append(visited, vline)
	}

	lines = draw(
		lines, func(curr pos, char rune) rune {
			if !slices.ContainsFunc(route, func(p pos) bool {
				return eq(p, curr)
			}) {
				return '.'
			}
			return char
		},
	)

	corner := route[0]
	for _, pipe := range route {
		if corner.line > pipe.line {
			corner = pipe
		} else if corner.line == pipe.line && corner.char > pipe.char {
			corner = pipe
		}
	}

	inside := pos{0, 1}
	mark(lines, corner, inside, inside, visited)

	inside = pos{1, 0}
	prev := add(corner, inside)
	mark(lines, corner, inside, prev, visited)

	total := 0
	for _, line := range lines {
		for _, char := range line {
			if char == 'I' {
				total += 1
			}
		}
	}
	fmt.Printf("%d\n", total)
}
