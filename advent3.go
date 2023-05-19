package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const file = "inputs.txt"

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	sum := 0
	sumBadge := 0
	i := 1
	var items []string

	for scanner.Scan() {
		line := scanner.Text()

		// Part 1
		items1, items2 := line[:len(line)/2], line[len(line)/2:]

		for _, item := range items1 {
			if strings.ContainsRune(items2, item) {
				sum += mapLetterToInt(item)
				break
			}
		}

		//Part 2
		items = append(items, line)
		if i%3 == 0 {
			sumBadge += calculateBadge(items)
			items = []string{}
		}
		i++
	}
	log.Println("Sum:", sum)
	log.Println("SumBadges", sumBadge)
}

func mapLetterToInt(letter rune) int {
	if letter >= 'a' && letter <= 'z' {
		return int(letter - 'a' + 1)
	} else {
		return int(letter - 'A' + 27)
	}
}

func calculateBadge(input []string) int {
	input0, input1, input2 := []rune(input[0]), input[1], input[2]
	for _, item := range input0 {
		if strings.ContainsRune(input1, item) && strings.ContainsRune(input2, item) {
			return mapLetterToInt(item)
		}
	}
	return 0
}
