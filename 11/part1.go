package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type pos struct {
	lnum int
	cnum int
}

func filter_range(values []int, start int, end int) []int {
	var filtered []int
	for _, value := range values {
		if value >= start && value <= end {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func get_dist(elines []int, ecols []int, first pos, second pos) int {
	lstart := int(math.Min(float64(first.lnum), float64(second.lnum)))
	lend := int(math.Max(float64(first.lnum), float64(second.lnum)))

	cstart := int(math.Min(float64(first.cnum), float64(second.cnum)))
	cend := int(math.Max(float64(first.cnum), float64(second.cnum)))

	ldist := (lend - lstart) + (cend - cstart) + len(filter_range(elines, lstart, lend)) + len(filter_range(ecols, cstart, cend))
	return ldist
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

	var elines []int
	var galaxies []pos
	for lnum, line := range lines {
		empty := true
		for cnum, char := range line {
			if char == '#' {
				galaxies = append(galaxies, pos{lnum, cnum})
				empty = false
			}
		}
		if empty {
			elines = append(elines, lnum)
		}
	}

	var ecols []int
	for cnum := range lines[0] {
		empty := true
		for _, line := range lines {
			char := line[cnum]
			if char == '#' {
				empty = false
				break
			}
		}
		if empty {
			ecols = append(ecols, cnum)
		}
	}

	// total := get_dist(elines, ecols, pos{5, 1}, pos{9, 4})
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ldist := get_dist(elines, ecols, galaxies[i], galaxies[j])
			total += ldist
		}
	}

	fmt.Printf("%v %v %d\n", elines, ecols, total)
}
