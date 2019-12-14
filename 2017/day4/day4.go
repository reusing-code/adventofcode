package day4

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func checkPhrasesFromFile(fileName string) int {
	res := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if checkPhrase(line) {
			res++
		}
	}

	return res
}

func checkPhrase(phrase string) bool {
	words := strings.Split(phrase, " ")
	if len(words) == 0 {
		return false
	}

	wordMap := make(map[string]bool)

	for _, word := range words {
		_, exists := wordMap[word]
		if exists {
			return false
		}
		wordMap[word] = true
	}

	return true
}

func checkPhrasesAnagramFromFile(fileName string) int {
	res := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if checkPhraseAnagram(line) {
			res++
		}
	}

	return res
}

func checkPhraseAnagram(phrase string) bool {
	words := strings.Split(phrase, " ")
	if len(words) == 0 {
		return false
	}

	wordMap := make(map[string]bool)

	for _, w := range words {
		word := getSortedWord(w)
		_, exists := wordMap[word]
		if exists {
			return false
		}
		wordMap[word] = true
	}

	return true
}

func getSortedWord(word string) string {
	b := []byte(word)
	sort.Sort(ascByte(b))
	return string(b)
}

type ascByte []byte

func (a ascByte) Len() int           { return len(a) }
func (a ascByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ascByte) Less(i, j int) bool { return a[i] < a[j] }
