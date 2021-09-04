package frequency

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strings"
)

//regex
//https://regexr.com

const expectedOutLength int = 10

func Frequency(s string) []string {
	if s == "" {
		return nil
	}

	s = strings.ToLower(s)
	splitted := strings.Split(s, " ")

	wordsSet := map[string]bool{}

	for _, word := range splitted {
		wordsSet[word] = true
	}

	//anonymous struct
	type wordCountPair struct {
		word  string
		count int
	}
	pairs := []wordCountPair{}
	for word := range wordsSet {
		if word == "" {
			continue
		}

		if regex, err := regexp.Compile(`\b` + word + `\b`); err != nil {
			log.Fatalf("regex for word '%s' can not be compiled", word)
		} else {
			result := regex.FindAll([]byte(s), -1)
			pairs = append(pairs, wordCountPair{word, len(result)})
		}
	}

	//sort by name to overcome problems with maps (order in which maps are iterated is always different)
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].word > pairs[j].word
	})
	//sort to finx X most frequent words
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	length := int(math.Min(float64(expectedOutLength), float64(len(pairs))))
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = pairs[i].word
	}

	return result
}
