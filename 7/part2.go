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

func has_count(occ map[rune]int, count int) int {
	// jokers := occ['J']
	has := 0
	for _, num := range occ {
		if num == count {
			has += 1
		}
		// if count > num && count-num <= jokers {
		// 	return true
		// }
	}
	return has
}

func get_rank(cards string) int {
	occ := make(map[rune]int)
	for _, c := range cards {
		if _, ok := occ[c]; !ok {
			occ[c] = 0
		}
		occ[c] += 1
	}

	jokers := occ['J']
	occ['J'] = 0

	most := 'J'
	for card, count := range occ {
		if count > occ[most] {
			most = card
		}
	}
	occ[most] += jokers

	if has_count(occ, 5) == 1 {
		return 7
	}
	if has_count(occ, 4) == 1 {
		return 6
	}
	if has_count(occ, 3) == 1 {
		if has_count(occ, 2) == 1 {
			return 5
		}
		return 4
	}

	if has_count(occ, 2) == 2 {
		return 3
	}

	if has_count(occ, 2) == 1 {
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
		return 0
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
	for _, line := range lines {
		parts := strings.Fields(line)
		bid, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
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
	for _, hand := range hands {
		fmt.Printf("%s %d\n", hand.cards, get_rank(hand.cards))
	}
	winnings := 0
	for rank, hand := range hands {
		winnings += (rank + 1) * hand.bid
	}
	fmt.Printf("%d\n", winnings)
	fmt.Printf("%d %d\n", len(lines), len(hands))

}
