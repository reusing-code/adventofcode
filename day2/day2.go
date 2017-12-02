package day2

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func checksum(fileName string) int {
	res := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text();
		values := strings.Split(line, "\t")
		if len(values) > 0 {
			min, _ := strconv.Atoi(values[0])
			max, _ := strconv.Atoi(values[0])
			for _, v := range values {
				intVal, _ := strconv.Atoi(v)
				if intVal > max {
					max = intVal
				}
				if intVal < min {
					min = intVal
				}
			}
			res += max -min
		}

	}

	return res
}

func checksumEvenDivisable(fileName string) int {
	res := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text();
		values := strings.Split(line, "\t")
		intArr := make([]int, 0, len(values))
		for _, v := range values {
			intVal, _ := strconv.Atoi(v)
			intArr = append(intArr, intVal)
		}

		res += evenDivisibleResult(intArr)

	}

	return res
}

func evenDivisibleResult(intArr []int) int {
	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			a := intArr[i]
			b := intArr[j]
			if a%b == 0 {
				return a / b
			}
			if b%a == 0 {
				return b / a
			}
		}
	}
	return 0
}
