package test

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"sort"

	"github.com/taz03/monkeytui/config"
)

const (
	BASE_URL   = "https://monkeytype.com/languages/"
	INIT_WORDS = 100
)

type languageModel struct {
	Name               string   `json:"name"`
	Ligatures          bool     `json:"ligatures"`
	NoLazyMode         bool     `json:"noLazyMode"`
	OrderedByFrequency bool     `json:"orderedByFrequency"`
	Words              []string `json:"words"`
}

func GenerateWords(config *config.Model, pulse chan bool) *[]string {
	language := fetchLanguage(config.Language)
	nextWord := language.nextWordFunc()

	switch config.Mode {
	case "time":
		return generateTimeTest(nextWord, pulse)
	case "words":
		return generateWordTest(nextWord, pulse, config.Words)
	}

	return &[]string{}
}

func fetchLanguage(name string) (language languageModel) {
	response, err := http.Get(BASE_URL + name + ".json")
	if err != nil || response.StatusCode != 200 {
		panic("Failed to fetch language")
	}

	defer response.Body.Close()

	bodySlice, _ := io.ReadAll(response.Body)
	json.Unmarshal(bodySlice, &language)

	return language
}

func (language *languageModel) nextWordFunc() func() string {
	n := len(language.Words)

	if language.OrderedByFrequency {
		var totalWeight int
		cumulativeWeights := make([]int, n)

		for i := 0; i < n; i++ {
			weight := (i + 1) * (n - i)
			totalWeight += weight
			if i == 0 {
				cumulativeWeights[i] = weight
			} else {
				cumulativeWeights[i] = cumulativeWeights[i-1] + weight
			}
		}

		return func() string {
			r := rand.Intn(totalWeight)
			// Use binary search to find the index of the chosen word
			idx := sort.Search(n, func(i int) bool { return cumulativeWeights[i] > r })
			return language.Words[idx]
		}
	}

	return func() string {
		return language.Words[rand.Intn(n)]
	}
}

func generateTimeTest(nextWord func() string, pulse chan bool) *[]string {
	generated := initialGenerator(INIT_WORDS, nextWord)

	go continiousGenerator(pulse, generated, nextWord)
	return generated
}

func generateWordTest(nextWord func() string, pulse chan bool, words int) *[]string {
	wordsToGenerate := words
	if words == 0 {
		wordsToGenerate = INIT_WORDS
	}

	generated := initialGenerator(wordsToGenerate, nextWord)

	if words == 0 {
		go continiousGenerator(pulse, generated, nextWord)
	}

	return generated
}

func initialGenerator(words int, nextWord func() string) *[]string {
	var generated []string

	for range words {
		generated = append(generated, nextWord())
	}

	return &generated
}

func continiousGenerator(pulse chan bool, generated *[]string, nextWord func() string) {
	for range pulse {
		*generated = append(*generated, nextWord())
	}
}
