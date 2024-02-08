package main

import (
	"fmt"
)

func main() {
	sentence := "hello world"
	counts := make(map[rune]int)

	for _, char := range sentence {
		counts[char]++
	}

	totalChars := 0
	for _, count := range counts {
		totalChars += count
	}

	for char, count := range counts {
		percentage := float64(count) / float64(totalChars) * 100
		fmt.Printf("%c - %d %.2f%%\n", char, count, percentage)
	}
}
