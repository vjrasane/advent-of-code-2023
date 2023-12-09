package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func chunk_slice(slice []int, size int) [][]int {
	var chunks [][]int
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < size {
			size = len(slice)
		}

		chunks = append(chunks, slice[0:size])
		slice = slice[size:]
	}

	return chunks
}

func to_nums(values []string) []int {
	var nums []int
	for _, value := range values {
		num, _ := strconv.Atoi(strings.TrimSpace(value))

		nums = append(nums, num)
	}
	return nums
}

func get_destination_range(nums []int, vrange []int) ([][]int, [][]int) {
	dstart := nums[0]
	length := nums[2]
	sstart := nums[1]
	send := sstart + length - 1

	vstart := vrange[0]
	vend := vrange[1]

	if vstart > vend {
		return make([][]int, 0), make([][]int, 0)
	}

	if vstart > send || vend < sstart {
		return make([][]int, 0), [][]int{vrange}
	}

	mstart := int(math.Max(float64(sstart), float64(vstart)))
	mend := int(math.Min(float64(send), float64(vend)))

	mapped := [][]int{{dstart + (mstart - sstart), dstart + (mend - sstart)}}

	var unmapped [][]int
	if sstart > vstart {
		unmapped = append(unmapped, []int{vstart, sstart - 1})
	}

	if send < vend {
		unmapped = append(unmapped, []int{send + 1, vend})
	}

	return mapped, unmapped
}

func get_destination_ranges(ranges [][]int, vrange []int) [][]int {
	var mapped [][]int

	slices := [][]int{vrange}

	for _, nums := range ranges {

		var unmapped [][]int
		for _, slice := range slices {
			_mapped, _unmapped := get_destination_range(nums, slice)
			mapped = append(mapped, _mapped...)
			unmapped = append(unmapped, _unmapped...)
		}
		slices = unmapped

	}
	mapped = append(mapped, slices...)
	return mapped
}

func get_destination_values(ranges [][]int, vranges [][]int) [][]int {
	var mapped [][]int
	for _, vrange := range vranges {
		_mapped := get_destination_ranges(ranges, vrange)
		mapped = append(mapped, _mapped...)
	}
	return mapped
}

func traverse(mappings map[string]map[string][][]int, source string, vranges [][]int) [][]int {
	lookup := mappings[source]

	var destination string
	for s := range lookup {
		destination = s
		break
	}

	ranges := lookup[destination]
	dvalues := get_destination_values(ranges, vranges)

	if destination == "location" {
		return dvalues
	}
	return traverse(mappings, destination, dvalues)
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

	seeds := chunk_slice(to_nums(strings.Fields(strings.Split(lines[0], ":")[1])), 2)
	lines = lines[1:len(lines)]

	var sections [][]string

	var section []string
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			section = append(section, line)
			continue
		}

		if len(section) > 0 {
			sections = append(sections, section)
		}
		section = make([]string, 0)
	}
	if len(section) > 0 {
		sections = append(sections, section)
	}

	mappings := make(map[string]map[string][][]int)
	for _, section := range sections {
		mapping := strings.Split(strings.Fields(section[0])[0], "-")
		source := mapping[0]
		destination := mapping[2]
		if _, ok := mappings[source]; !ok {
			mappings[source] = make(map[string][][]int)
		}

		lines := section[1:]
		var ranges [][]int
		for _, line := range lines {
			nums := to_nums(strings.Fields(line))
			ranges = append(ranges, nums)
		}

		mappings[source][destination] = ranges
	}

	var vranges [][]int
	for _, seed := range seeds {
		start := seed[0]
		length := seed[1]
		end := start + length - 1
		vranges = append(vranges, []int{start, end})

	}
	lranges := traverse(mappings, "seed", vranges)

	lowest := lranges[0][0]
	for _, lrange := range lranges[1:] {
		if lowest > lrange[0] {
			lowest = lrange[0]
		}
	}
	fmt.Printf("%v %d\n", lranges, lowest)
}
