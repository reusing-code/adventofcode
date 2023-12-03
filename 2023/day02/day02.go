package main

import (
	"fmt"
	"strings"
)

var colorMap = map[string]int{
	"red":   0,
	"green": 1,
	"blue":  2,
}

type round struct {
	colors [3]int
}

type game struct {
	id     int
	rounds []*round
}

func parseGames(input []string) []*game {
	games := make([]*game, 0)
	for _, line := range input {
		game := &game{}
		fmt.Sscanf(line, "Game %d:", &game.id)
		sep := strings.Split(line, ":")
		sep = strings.Split(sep[1], ";")
		rounds := make([]*round, 0)
		for _, r := range sep {
			round := &round{}
			for _, s := range strings.Split(r, ",") {
				var val int
				var color string
				fmt.Sscanf(s, "%d %s", &val, &color)
				round.colors[colorMap[color]] = val
			}
			rounds = append(rounds, round)
		}
		game.rounds = rounds
		games = append(games, game)

	}
	return games
}

func puzzle1(input []string) int {
	games := parseGames(input)
	result := 0
	max := &round{[3]int{12, 13, 14}}
	for _, game := range games {
		possible := true
		for _, r := range game.rounds {
			for i, v := range max.colors {
				if r.colors[i] > v {
					possible = false
				}
			}
		}
		if possible {
			result += game.id
		}
	}
	return result
}

func puzzle2(input []string) int {
	games := parseGames(input)
	result := 0
	for _, game := range games {
		max := &round{}
		for _, r := range game.rounds {
			for i, v := range r.colors {
				if v > max.colors[i] {
					max.colors[i] = v
				}
			}
		}
		result += max.colors[0] * max.colors[1] * max.colors[2]
	}
	return result
}
