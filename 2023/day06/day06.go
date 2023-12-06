package main

import (
	"fmt"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func calcDistanceReached(time, hold int) int {
	return hold * (time - hold)
}

func binarySeachAsc(time, distance, from, to int) int {
	if from == to {
		fmt.Println(time, distance, from, to)
		panic("Asc from == to not allowed")
	}
	i := (to + from) / 2
	iDist := calcDistanceReached(time, i)
	i1Dist := calcDistanceReached(time, i+1)
	if iDist <= distance && i1Dist > distance {
		return i + 1
	}
	if iDist >= i1Dist {
		return binarySeachAsc(time, distance, from, i)
	}
	if iDist <= distance {
		return binarySeachAsc(time, distance, i, to)
	} else {
		return binarySeachAsc(time, distance, from, i)
	}
}

func binarySeachDesc(time, distance, from, to int) int {
	if from == to {
		fmt.Println(time, distance, from, to)
		panic("Desc from == to not allowed")
	}
	i := (to + from) / 2
	iDist := calcDistanceReached(time, i)
	i1Dist := calcDistanceReached(time, i+1)
	if iDist > distance && i1Dist <= distance {
		return i
	}
	if iDist <= i1Dist {
		return binarySeachDesc(time, distance, i, to)
	}
	if iDist > distance {
		return binarySeachDesc(time, distance, i, to)
	} else {
		return binarySeachDesc(time, distance, from, i)
	}
}

func puzzle1(input []string) int {
	timeStrs := strings.Fields(input[0])
	distanceStrs := strings.Fields(input[1])
	races := make([]race, 0)
	for i := 1; i < len(timeStrs); i++ {
		time, _ := strconv.Atoi(timeStrs[i])
		distance, _ := strconv.Atoi(distanceStrs[i])
		races = append(races, race{time, distance})
	}
	result := 1
	for _, race := range races {
		result *= binarySeachDesc(race.time, race.distance, 0, race.time) - binarySeachAsc(race.time, race.distance, 0, race.time) + 1
	}
	return result
}

func puzzle2(input []string) int {
	var time int
	var distance int
	fmt.Sscanf(strings.ReplaceAll(input[0], " ", ""), "Time:%d", &time)
	fmt.Sscanf(strings.ReplaceAll(input[1], " ", ""), "Distance:%d", &distance)
	result := binarySeachDesc(time, distance, 0, time) - binarySeachAsc(time, distance, 0, time) + 1
	return result
}
