package test

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/taz03/monkeytui/config"
)

const BASE_URL = "https://monkeytype.com/languages/"

type languageModel struct {
	Name               string   `json:"name"`
    Ligatures          bool     `json:"ligatures"`
	NoLazyMode         bool     `json:"noLazyMode"`
	OrderedByFrequency bool     `json:"orderedByFrequency"`
	Words              []string `json:"words"`
}

func GenerateWords(config *config.Model) *[]string {
    language := fetchLanguage(config.Language)

	switch config.Mode {
	case "time":
        return language.generateInfiniteTest()
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

func (this *languageModel) generateInfiniteTest() *[]string {
	n := len(this.Words)

    totalWeight := 0
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

    var generated []string

    appendRandomWord := func() {
        r := rand.Intn(totalWeight)
        // Use binary search to find the index of the chosen word
        idx := sort.Search(n, func(i int) bool { return cumulativeWeights[i] > r })
        generated = append(generated, this.Words[idx])
    }

    for range 50 {
        appendRandomWord()
    }

    var delay time.Duration = 100
    go func() {
        for {
            appendRandomWord()
            delay += 10
            time.Sleep(delay * time.Millisecond)
        }
    }()

    return &generated
}
