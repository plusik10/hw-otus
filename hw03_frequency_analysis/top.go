package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func Top10(str string) []string {
	if str == "" {
		return []string{}
	}

	words := strings.Fields(str)
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}
	wordsSlice := make([]WordCount, 0, len(wordsMap))

	for word, count := range wordsMap {
		wordsSlice = append(wordsSlice, WordCount{word, count})
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].Count == wordsSlice[j].Count {
			return wordsSlice[i].Word < wordsSlice[j].Word
		}
		return wordsSlice[i].Count > wordsSlice[j].Count
	})

	return WordsToStringSlice(wordsSlice)
}

func WordsToStringSlice(words []WordCount) []string {
	result := make([]string, 0, len(words))
	for _, word := range words {
		result = append(result, word.Word)
	}

	if len(result) > 10 {
		return result[:10]
	}
	return result
}
