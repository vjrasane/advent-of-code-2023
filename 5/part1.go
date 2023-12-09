package old

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

func get_destintation_value(ranges [][]int, value int) int {
	for _, nums := range ranges {
		dstart := nums[0]
		sstart := nums[1]
		length := nums[2]

		fmt.Printf("%d %d %d\n", dstart, sstart, length)

		if value >= sstart && value <= sstart+length-1 {

			fmt.Println("MAPPED")
			return dstart + (value - sstart)
		}
	}
	return value
}

func traverse(mappings map[string]map[string][][]int, source string, value int) int {
	lookup := mappings[source]

	var destination string
	for s := range lookup {
		destination = s
		break
	}

	ranges := lookup[destination]
	dvalue := get_destintation_value(ranges, value)

	// fmt.Printf("%s %d %s %d\n", source, value, destination, value)
	if destination == "location" {
		return dvalue
	}
	return traverse(mappings, destination, dvalue)
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

	seeds := to_nums(strings.Fields(strings.Split(lines[0], ":")[1]))
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

		// fmt.Printf("%s %s %v\n", source, destination, ranges)
	}

	var locations []int
	for _, seed := range seeds {
		location := traverse(mappings, "seed", seed)
		locations = append(locations, location)
	}

	lowest := locations[0]
	for _, loc := range locations[1:] {
		if lowest > loc {
			lowest = loc
		}
	}
	fmt.Printf("%v %v %d\n", seeds, locations, lowest)
}
