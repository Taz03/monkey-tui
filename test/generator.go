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
    BASE_URL = "https://monkeytype.com/languages/"
    WORDS_PER_SECOND = 50
)

type languageModel struct {
	Name               string   `json:"name"`
    Ligatures          bool     `json:"ligatures"`
	NoLazyMode         bool     `json:"noLazyMode"`
	OrderedByFrequency bool     `json:"orderedByFrequency"`
	Words              []string `json:"words"`
}

func GenerateWords(config *config.Model) []string {
    language := fetchLanguage(config.Language)

    var nextWord func() string
    if language.OrderedByFrequency {
        n := len(language.Words)

        totalWeight, cumulativeWeights := calculateWeight(n)
        nextWord = func() string {
            r := rand.Intn(totalWeight)
            // Use binary search to find the index of the chosen word
            idx := sort.Search(n, func(i int) bool { return cumulativeWeights[i] > r })
            return language.Words[idx]
        }
    } else {
        nextWord = func() string {
            return language.Words[rand.Intn(len(language.Words))]
        }
    }

	switch config.Mode {
	case "time":
        return language.generateTimeTest(config.Time, nextWord)
	}

	return []string{}
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

func (this *languageModel) generateTimeTest(seconds int, nextWord func() string) []string {
    var generated []string
    for range seconds * WORDS_PER_SECOND {
        generated = append(generated, nextWord())
    }

    return generated
}

func calculateWeight(n int) (totalWeight int, cumulativeWeights []int) {
	cumulativeWeights = make([]int, n)
	for i := 0; i < n; i++ {
		weight := (i + 1) * (n - i)
		totalWeight += weight
		if i == 0 {
			cumulativeWeights[i] = weight
		} else {
			cumulativeWeights[i] = cumulativeWeights[i-1] + weight
		}
	}

    return
}
