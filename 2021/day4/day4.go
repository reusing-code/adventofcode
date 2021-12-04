package main

import (
	"github.com/reusing-code/adventofcode/gohelpers"
	"strings"
	"strconv"
)

type board struct {
	numbers [5][5]int
	marked [5][5]bool
}

func puzzle1(input []string) int {
	// Parse input
	split := gohelpers.SplitByEmptyLine(input)
	moves := strings.Split(split[0][0], ",")
	boards := make([]*board, 0, len(split) - 1)
	for i, boardSlice := range split {
		if i == 0 {
			continue
		}
		boards = append(boards, parseBoard(boardSlice))
	}

	// Go through all drawn numbers
	for _, drawnNumber := range moves {
		number, err := strconv.Atoi(drawnNumber)
		if err != nil {
			panic(err)
		}
		for _, board := range boards {
			board.mark(number)
			if board.hasWon() {
				return number * board.calcScore()
			}
		}
	}

	return 0
}

func puzzle2(input []string) int {
		// Parse input
		split := gohelpers.SplitByEmptyLine(input)
		moves := strings.Split(split[0][0], ",")
		boards := make([]*board, 0, len(split) - 1)
		for i, boardSlice := range split {
			if i == 0 {
				continue
			}
			boards = append(boards, parseBoard(boardSlice))
		}
	
		winners := 0
	
		// Go through all drawn numbers
		for _, drawnNumber := range moves {
			number, err := strconv.Atoi(drawnNumber)
			if err != nil {
				panic(err)
			}
			for _, board := range boards {
				if board.hasWon() {
					continue
				}
				board.mark(number)
				if board.hasWon() {
					winners++
					if (winners == len(boards)) {
						return number * board.calcScore()
					}
				}
			}
		}
	
		return 0
}

func parseBoard(in []string) *board {
	result := &board{}
	for i, line := range in {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "  ", " ")
		split := strings.Split(line, " ")
		for j, str := range split {
			val, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			result.numbers[i][j] = val
		}
	}
	return result
}

func (b *board) mark(val int) {
	for i, line := range b.numbers {
		for j, number := range line {
			if number == val {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *board) hasWon() bool {
	// rows
	for i, line := range b.numbers {
		allMarked := true
		for j, _ := range line {
			if !b.marked[i][j] {
				 allMarked = false
				 break
			}
		}
		if allMarked {
			return true
		}
	}
	// columns
	for j := 0; j < 5; j++ {
		allMarked := true
		for i := 0; i < 5; i++ {
			if !b.marked[i][j] {
				 allMarked = false
				 break
			}
		}
		if allMarked {
			return true
		}
	}
	return false
}

func (b *board) calcScore() int {
	result := 0
	for i, line := range b.numbers {
		for j, number := range line {
			if !b.marked[i][j] {
				 result += number
			}
		}
	}
	return result
}
