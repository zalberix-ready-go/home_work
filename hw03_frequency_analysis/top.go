package hw03frequencyanalysis

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var collateRussian = collate.New(language.Russian)

func Top10(text string) []string {
	text = strings.ToLower(text)

	words := strings.Fields(text)

	words, err := normalazeWords(words)
	if err != nil {
		log.Panic("normalazed words: %w", err)

		return nil
	}

	frequencyWords := make(map[string]int)

	// Подсчет частоты слов
	for _, word := range words {
		// В случае если нет такого ключа то для int zero value - 0
		frequencyWords[word]++
	}

	wordsGroupByFrequency := make(map[int][]string)

	// Группировака слов по частоте появления
	for word, count := range frequencyWords {
		wordsGroupByFrequency[count] = append(wordsGroupByFrequency[count], word)
	}

	allGroup := make([]int, 0)

	// Получение всех групп (ключей)
	for group := range wordsGroupByFrequency {
		allGroup = append(allGroup, group)
	}

	// Сортируем по убыванию
	sort.Slice(allGroup, func(i, j int) bool {
		return allGroup[i] > allGroup[j]
	})

	topFrequencyWord := make([]string, 0)

	for _, group := range allGroup {
		collateRussian.SortStrings(wordsGroupByFrequency[group])

		topFrequencyWord = append(topFrequencyWord, wordsGroupByFrequency[group]...)

		if len(topFrequencyWord) >= 10 {
			break
		}
	}

	if len(topFrequencyWord) > 10 {
		return topFrequencyWord[:10]
	}

	return topFrequencyWord
}

// Убирает все знаки препинания по бокам

func normalazeWords(words []string) ([]string, error) {
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("произошла паника при компиляции регулярного выражения: %v", r)
		}
	}()

	re := regexp.MustCompile(`^[^а-яА-ЯёЁa-zA-Z0-9]+|[^а-яА-ЯёЁa-zA-Z0-9]+$`)

	normalWords := make([]string, 0, len(words))

	for _, word := range words {
		if word == "-" {
			continue
		}
		normalWord := re.ReplaceAllString(word, "")
		if normalWord != "" {
			normalWords = append(normalWords, normalWord)
		} else {
			normalWords = append(normalWords, word)
		}
	}
	return normalWords, err
}
