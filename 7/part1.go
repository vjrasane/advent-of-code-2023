package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type hand struct {
	cards string
	bid   int
}

func has_count(occ map[rune]int, count int) bool {
	for _, num := range occ {
		if num == count {
			return true
		}
	}
	return false
}

func get_rank(cards string) int {
	occ := make(map[rune]int)
	for _, c := range cards {
		if _, ok := occ[c]; !ok {
			occ[c] = 0
		}
		occ[c] += 1
	}

	if has_count(occ, 5) {
		return 7
	}
	if has_count(occ, 4) {
		return 6
	}
	if has_count(occ, 3) {
		if has_count(occ, 2) {
			return 5
		}
		return 4
	}

	if has_count(occ, 2) {
		if len(occ) == 3 {
			return 3
		}
		return 2
	}
	return 1
}

func card_to_int(card rune) int {
	if unicode.IsDigit(card) {
		num, _ := strconv.Atoi(string(card))
		return num
	}
	if card == 'A' {
		return 14
	}
	if card == 'K' {
		return 13
	}
	if card == 'Q' {
		return 12
	}

	if card == 'J' {
		return 11
	}

	if card == 'T' {
		return 10
	}

	return 0
}

func get_order(cards string) int {
	total := 0
	for i, card := range cards {
		total += card_to_int(card) * int(math.Pow(10, float64(len(cards)-i-1)))
	}
	return total

}

func card_order(acards string, bcards string) int {
	if len(acards) == 0 {
		return 0
	}
	anum := card_to_int([]rune(acards)[0])
	bnum := card_to_int([]rune(bcards)[0])
	if anum == bnum {
		return card_order(acards[1:], bcards[1:])
	}
	return cmp.Compare(anum, bnum)

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

	var hands []hand
	// var bids []int
	for _, line := range lines {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, hand{parts[0], bid})
	}

	cmp_hands := func(a, b hand) int {
		arank := get_rank(a.cards)
		brank := get_rank(b.cards)
		if arank == brank {
			return card_order(a.cards, b.cards)
		}
		return cmp.Compare(arank, brank)
	}

	slices.SortFunc(hands, cmp_hands)

	winnings := 0
	for rank, hand := range hands {
		winnings += (rank + 1) * hand.bid
	}
	fmt.Printf("%v %d", hands, winnings)
}
