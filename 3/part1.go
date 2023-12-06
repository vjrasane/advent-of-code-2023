package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func get_num_at(lines []string, linenum int, char int) (*int, *int) {
	line := lines[linenum]
	beginning := regexp.MustCompile(`[0-9]*$`).FindStringSubmatch(line[0 : char+1])[0]
	ending := regexp.MustCompile(`^[0-9]*`).FindStringSubmatch(line[char:len(line)])[0]
	if len(beginning) > 0 {
		beginning = beginning[0 : len(beginning)-1]
	}
	num, err := strconv.Atoi(beginning + ending)
	if err != nil {
		return nil, nil
	}

	index := char - len(beginning)
	return &num, &index
}

func has_num(nums [][]int, num int, linenum int, charnum int) bool {
	for _, found := range nums {
		if found[0] == num && found[1] == linenum && found[2] == charnum {
			return true
		}
	}
	return false
}

func get_surrounding(lines []string, linenum int, char int) []int {
	dirs := [][]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}

	var nums [][]int
	for _, dir := range dirs {
		linenum_at := linenum + dir[1]
		char_at := char + dir[0]

		if linenum_at >= len(lines) || char_at > len(lines[linenum_at]) {
			continue
		}

		num, index := get_num_at(lines, linenum_at, char_at)
		if num != nil && !has_num(nums, *num, linenum_at, *index) {
			nums = append(nums, []int{*num, linenum_at, *index})
		}
	}

	var values []int
	for _, num := range nums {
		values = append(values, num[0])
	}
	return values
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

	total := 0
	for linenum, line := range lines {
		for charnum, c := range line {
			if c == '.' || unicode.IsDigit(c) {
				continue
			}
			nums := get_surrounding(lines, linenum, charnum)
			for _, num := range nums {
				total += num
			}
		}
	}
	fmt.Println(total)

}
