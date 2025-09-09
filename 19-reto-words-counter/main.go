package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanText(text string) string {
	regex1 := regexp.MustCompile(`-`)
	cleanText := regex1.ReplaceAllString(text, " ")
	
	regex2 := regexp.MustCompile(`[^a-zA-Z0-9\s']`)
	cleanText = regex2.ReplaceAllString(cleanText, " ")
	cleanText = strings.ToLower(cleanText)

	return cleanText
}

func cleanApostofres(word string) string {
	regex := regexp.MustCompile(`'`)
	cleanWord := regex.ReplaceAllString(word, "")

	return cleanWord
}

func CountWordFrequency(text string) map[string]int {
	cleanText := cleanText(text)
	
	words := strings.Fields(cleanText)
	
	wordFrequency := make(map[string]int)

	fmt.Println("Texto limpio:", cleanText)
	fmt.Println("Palabras:", words)

	for _, word := range words {
		cleanWord := cleanApostofres(word)
		wordFrequency[cleanWord] = wordFrequency[cleanWord] + 1
	}

	return wordFrequency
} 

func main() {
	text := "  Spaces,   tabs,\t\tand\nnew-lines are ignored!  "
	wordFrequency := CountWordFrequency(text)
	fmt.Println(wordFrequency)
}