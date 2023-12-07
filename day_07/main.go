package main

import (
	"bytes"
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"sort"
	"strconv"
)

type HandRank int

const (
	//12345
	HighCard HandRank = iota
	//12344
	OnePair
	//12233
	TwoPairs
	//12333
	ThreeOfAKind
	//11222
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []byte
	Rank  HandRank
	Bid   int
}

var CardStrengthNormal = map[byte]int{
	byte('2'): 0,
	byte('3'): 1,
	byte('4'): 2,
	byte('5'): 3,
	byte('6'): 5,
	byte('7'): 6,
	byte('8'): 7,
	byte('9'): 8,
	byte('T'): 9,
	byte('J'): 10,
	byte('Q'): 11,
	byte('K'): 12,
	byte('A'): 13,
}
var CardStrengthWithJoker = map[byte]int{
	byte('2'): 0,
	byte('3'): 1,
	byte('4'): 2,
	byte('5'): 3,
	byte('6'): 5,
	byte('7'): 6,
	byte('8'): 7,
	byte('9'): 8,
	byte('T'): 9,
	byte('J'): -1,
	byte('Q'): 11,
	byte('K'): 12,
	byte('A'): 13,
}

func (h *Hand) getRankWithoutJoker() (rank HandRank) {
	var cards = string(h.Cards)
	var cardMap = map[string]int{}
	for _, card := range cards {
		cardMap[string(card)] += 1
	}
	var cardCount = len(cardMap)
	var highestCount = 0
	for _, count := range cardMap {
		if count > highestCount {
			highestCount = count
		}
	}
	switch cardCount {
	case 1:
		rank = FiveOfAKind
	case 2:
		if highestCount == 4 {
			rank = FourOfAKind
		} else {
			rank = FullHouse
		}
	case 3:
		if highestCount == 3 {
			rank = ThreeOfAKind
		} else {
			rank = TwoPairs
		}
	case 4:
		rank = OnePair
	case 5:
		rank = HighCard
	}

	return
}
func (h *Hand) getRankWithJoker() (rank HandRank) {
	var cards = string(h.Cards)
	var cardMap = map[string]int{}
	for i, card := range cards {
		if string(card) == "J" {
			cardMap[string(card)+string(rune(i+1))] += 1
		}
		cardMap[string(card)] += 1
	}
	var jokerCount = cardMap["J"]
	delete(cardMap, "J")

	var cardCount = len(cardMap)
	var highestCount = 0
	for c, count := range cardMap {
		if c != "J" && count > highestCount {
			highestCount = count
		}
	}
	switch cardCount {
	case 1:
		rank = FiveOfAKind
	case 2:
		if highestCount == 4 {
			rank = FourOfAKind
		} else {
			rank = FullHouse
		}
	case 3:
		if highestCount == 3 {
			rank = ThreeOfAKind
		} else {
			rank = TwoPairs
		}
	case 4:
		rank = OnePair
	case 5:
		rank = HighCard
	}
	var jokerMap = map[HandRank]HandRank{
		HighCard:     OnePair,
		OnePair:      ThreeOfAKind,
		ThreeOfAKind: FourOfAKind,
		FourOfAKind:  FiveOfAKind,
		TwoPairs:     FullHouse,
	}
	if jokerCount == 5 {
		rank = FiveOfAKind
	} else {
		for i := 0; i < jokerCount; i++ {
			rank = jokerMap[rank]
		}
	}

	return
}

func (h *Hand) compare(other Hand, cardRank map[byte]int) int {
	if h.Rank > other.Rank {
		return 1
	} else if h.Rank < other.Rank {
		return -1
	}

	for i := 0; i < 5; i++ {
		if cardRank[h.Cards[i]] > cardRank[other.Cards[i]] {
			return 1
		} else if cardRank[h.Cards[i]] < cardRank[other.Cards[i]] {
			return -1
		}
	}
	return 0
}

func parseData(data [][]byte) []Hand {
	var hands []Hand
	for _, line := range data {
		var hand Hand
		var splittedLine = bytes.Split(line, []byte(" "))
		var cards, bid = splittedLine[0], splittedLine[1]

		hand.Cards = cards
		hand.Bid, _ = strconv.Atoi(string(bid))
		hands = append(hands, hand)
	}
	return hands
}

func partOne(hands []Hand) (out int) {
	for i, _ := range hands {
		hands[i].Rank = hands[i].getRankWithoutJoker()
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].compare(hands[j], CardStrengthNormal) != 1
	})
	for i, hand := range hands {
		out += hand.Bid * (i + 1)
	}
	return
}
func partTwo(hands []Hand) (out int) {
	for i, _ := range hands {
		hands[i].Rank = hands[i].getRankWithJoker()
	}
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].compare(hands[j], CardStrengthWithJoker) != 1
	})
	for i, hand := range hands {
		out += hand.Bid * (i + 1)
	}
	return
}

func main() {
	var data = utils.GetInput("day_07/input.txt")
	var hands = parseData(data)

	fmt.Println("Part one:", partOne(hands))
	fmt.Println("Part Two:", partTwo(hands))
}
