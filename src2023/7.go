package src2023

import (
	"sort"
	"strconv"
	"strings"

	"github.com/vkalekis/advent-of-code/utils"
	"go.uber.org/zap"
)

var cardTypes []rune

type hand string

type round struct {
	h   hand
	wt  wintype
	bid int
}

type wintype int64

const (
	fiveofakind wintype = iota
	fourofakind
	fullhouse
	threeofakind
	twopair
	onepair
	high
)

func (r *round) parseWinType(part int, logger *zap.SugaredLogger) {
	cardsCount := make(map[string]int)

	for _, card := range r.h {
		if _, ok := cardsCount[string(card)]; ok {
			cardsCount[string(card)]++
		} else {
			cardsCount[string(card)] = 1
		}
	}

	var wt wintype

	twopairscount := 0
	threepairscount := 0

	logger.Debugf("CardCount: %v", cardsCount)

	// default - worse win scenario
	wt = high
	jokers := 0

	for card, count := range cardsCount {
		if card == "J" && part == 2 {
			jokers = count
			continue
		}

		switch count {
		case 5:
			wt = fiveofakind
		case 4:
			wt = fourofakind
		case 3:
			threepairscount++
		case 2:
			twopairscount++
		}

		if threepairscount == 1 {
			if twopairscount == 1 {
				wt = fullhouse
			} else {
				wt = threeofakind
			}
		}
		if twopairscount == 2 {
			wt = twopair
		}
		if twopairscount == 1 && threepairscount == 0 {
			wt = onepair
		}

	}

	r.wt = wt

	if jokers > 0 && part == 2 {
		for j := 0; j < jokers; j++ {
			r.addJoker(logger)
		}
	}

	logger.Infof("Hand: %v Cards: %+v Win: %v", r.h, cardsCount, r.wt)

}

func (r *round) addJoker(logger *zap.SugaredLogger) {
	var newwt wintype
	switch r.wt {
	case high:
		newwt = onepair
	case onepair:
		newwt = threeofakind
	case twopair:
		newwt = fullhouse
	case threeofakind:
		newwt = fourofakind
	case fourofakind:
		newwt = fiveofakind
	}

	logger.Debugf("Joker: Hand: %v Old wt: %v New wt: %v", r.h, r.wt, newwt)
	r.wt = newwt
}

func findIndex(arr []rune, target rune) int {
	for ind, v := range arr {
		if v == target {
			return ind
		}
	}
	return -1
}

func (h1 hand) compareHands(h2 hand) bool {
	runes2 := []rune(h2)
	for ind, card1 := range h1 {

		cardtype1 := findIndex(cardTypes, card1)
		cardtype2 := findIndex(cardTypes, runes2[ind])

		res := cardtype1 - cardtype2
		if res == 0 {
			continue
		} else if res > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func Day_07(part int, reader utils.Reader, logger *zap.SugaredLogger) int {

	switch part {
	case 1:
		cardTypes = []rune("AKQJT98765432")
	case 2:
		cardTypes = []rune("AKQT98765432J")
	}

	rounds := make([]round, 0)

	for line := range reader.Stream() {
		splitLine := strings.Split(utils.StandardizeSpaces(line), " ")
		bid, err := strconv.Atoi(splitLine[1])
		if err != nil {
			logger.Errorf("Error while parsing int: %v", err)
			return -1
		}

		r := round{
			h:   hand(splitLine[0]),
			bid: bid,
		}

		r.parseWinType(part, logger)

		rounds = append(rounds, r)
	}

	logger.Debugf("Rounds: %+v", rounds)

	// collect the results on a map with the wintype as the key
	roundsMap := make(map[wintype][]round)

	for _, r := range rounds {
		if _, ok := roundsMap[r.wt]; !ok {
			roundsMap[r.wt] = make([]round, 0)
		}
		roundsMap[r.wt] = append(roundsMap[r.wt], r)
	}

	// loop through all winttype scenarios in reverse order (worst to best)
	// assigning ranks starting from 1 (worse hand) -> .. (best hand)
	// then calculate the total winnings by adding a product of bid*rank for each hand
	wts := []wintype{high, onepair, twopair, threeofakind, fullhouse, fourofakind, fiveofakind}
	rank := 1
	totalWinnings := 0

	logger.Debugf("Rounds: %+v", roundsMap[6])

	for _, wt := range wts {
		if _, ok := roundsMap[wt]; !ok {
			continue
		}

		if len(roundsMap[wt]) == 1 {
			logger.Debugf("%d * %d = %d", rank, roundsMap[wt][0].bid, rank*roundsMap[wt][0].bid)

			totalWinnings += rank * roundsMap[wt][0].bid
			rank++
		} else {
			// in the case of multiple hands per win type, sort by comparing the two hands
			sort.Slice(roundsMap[wt], func(i, j int) bool {
				return roundsMap[wt][i].h.compareHands(roundsMap[wt][j].h)
			})

			logger.Debugf("Rounds: %+v", roundsMap[wt])

			for _, r := range roundsMap[wt] {
				logger.Debugf("WT: %v : %d * %d = %d", wt, rank, r.bid, rank*r.bid)
				totalWinnings += rank * r.bid
				rank++
			}
		}
	}

	return totalWinnings
}
