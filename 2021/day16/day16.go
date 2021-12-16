package main

import (
	"math"
	"strconv"
)

var mapping = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func arithmetic(typeId int, params []int) int {
	if typeId == 0 {
		result := 0
		for _, p := range params {
			result += p
		}
		return result
	}
	if typeId == 1 {
		result := 1
		for _, p := range params {
			result *= p
		}
		return result
	}
	if typeId == 2 {
		result := math.MaxInt
		for _, p := range params {
			if p < result {
				result = p
			}
		}
		return result
	}
	if typeId == 3 {
		result := 0
		for _, p := range params {
			if p > result {
				result = p
			}
		}
		return result
	}
	if typeId == 5 {
		if params[0] > params[1] {
			return 1
		} else {
			return 0
		}
	}
	if typeId == 6 {
		if params[0] < params[1] {
			return 1
		} else {
			return 0
		}
	}
	if typeId == 7 {
		if params[0] == params[1] {
			return 1
		} else {
			return 0
		}
	}
	return 1
}

func puzzle1(input []string) int {
	bitArray := ""
	for _, r := range input[0] {
		bitArray = bitArray + mapping[r]
	}
	version := 0
	var parsePacket func(int) (int, int)
	parsePacket = func(idx int) (val, consumed int) {
		ver, err := strconv.ParseInt(string(bitArray[idx:idx+3]), 2, 8)
		if err != nil {
			panic(err)
		}
		version += int(ver)
		typeID, err := strconv.ParseInt(string(bitArray[idx+3:idx+6]), 2, 8)
		if err != nil {
			panic(err)
		}
		consumed = 6
		if typeID == 4 {
			literal := 0
			for {
				literal *= 16
				add, err := strconv.ParseInt(string(bitArray[idx+consumed+1:idx+consumed+5]), 2, 8)
				if err != nil {
					panic(err)
				}
				literal += int(add)
				consumed += 5
				if bitArray[idx+consumed-5] == '0' {
					break
				}
			}
			val = literal
		} else {
			lengthTypeId := bitArray[idx+consumed]
			consumed++
			params := make([]int, 0)
			if lengthTypeId == '0' {
				bitlength, err := strconv.ParseInt(string(bitArray[idx+consumed:idx+consumed+15]), 2, 16)
				if err != nil {
					panic(err)
				}
				consumed += 15
				end := consumed + int(bitlength)

				for consumed < end {
					v, cons := parsePacket(idx + consumed)
					consumed += cons
					params = append(params, v)
				}
			} else {
				subpackets, err := strconv.ParseInt(string(bitArray[idx+consumed:idx+consumed+11]), 2, 16)
				if err != nil {
					panic(err)
				}
				consumed += 11
				for i := 0; i < int(subpackets); i++ {
					v, cons := parsePacket(idx + consumed)
					consumed += cons
					params = append(params, v)
				}
			}
			val = arithmetic(int(typeID), params)
		}
		return
	}

	parsePacket(0)

	return version
}

func puzzle2(input []string) int {
	bitArray := ""
	for _, r := range input[0] {
		bitArray = bitArray + mapping[r]
	}
	version := 0
	var parsePacket func(int) (int, int)
	parsePacket = func(idx int) (val, consumed int) {
		ver, err := strconv.ParseInt(string(bitArray[idx:idx+3]), 2, 8)
		if err != nil {
			panic(err)
		}
		version += int(ver)
		typeID, err := strconv.ParseInt(string(bitArray[idx+3:idx+6]), 2, 8)
		if err != nil {
			panic(err)
		}
		consumed = 6
		if typeID == 4 {
			literal := 0
			for {
				literal *= 16
				add, err := strconv.ParseInt(string(bitArray[idx+consumed+1:idx+consumed+5]), 2, 8)
				if err != nil {
					panic(err)
				}
				literal += int(add)
				consumed += 5
				if bitArray[idx+consumed-5] == '0' {
					break
				}
			}
			val = literal
		} else {
			lengthTypeId := bitArray[idx+consumed]
			consumed++
			params := make([]int, 0)
			if lengthTypeId == '0' {
				bitlength, err := strconv.ParseInt(string(bitArray[idx+consumed:idx+consumed+15]), 2, 16)
				if err != nil {
					panic(err)
				}
				consumed += 15
				end := consumed + int(bitlength)

				for consumed < end {
					v, cons := parsePacket(idx + consumed)
					consumed += cons
					params = append(params, v)
				}
			} else {
				subpackets, err := strconv.ParseInt(string(bitArray[idx+consumed:idx+consumed+11]), 2, 16)
				if err != nil {
					panic(err)
				}
				consumed += 11
				for i := 0; i < int(subpackets); i++ {
					v, cons := parsePacket(idx + consumed)
					consumed += cons
					params = append(params, v)
				}
			}
			val = arithmetic(int(typeID), params)
		}
		return
	}

	val, _ := parsePacket(0)

	return val
}
