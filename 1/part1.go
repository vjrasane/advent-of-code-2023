package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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
		first := regexp.MustCompile(`^[a-zA-Z]*([0-9])`).FindStringSubmatch(line)[1]
		last := regexp.MustCompile(`([0-9])[a-zA-Z]*$`).FindStringSubmatch(line)[1]
		value, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		total += value
	}
	fmt.Println(total)
}
