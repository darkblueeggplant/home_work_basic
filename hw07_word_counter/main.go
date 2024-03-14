package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(text) //  функции для работы со строками, чтобы разделить текст на отдельные слова
	// println(len(words)) 7
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++ //
		// println(word)
	}
	return wordCounts
}

func main() {
	text := "рельсы рельсы шпалы шпалы едет поезд запоздалый"
	wordCountMap := countWords(text)
	fmt.Println(wordCountMap)
}
