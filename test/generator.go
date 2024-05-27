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
    INIT_WORDS = 100
)

type languageModel struct {
	Name               string   `json:"name"`
    Ligatures          bool     `json:"ligatures"`
	NoLazyMode         bool     `json:"noLazyMode"`
	OrderedByFrequency bool     `json:"orderedByFrequency"`
	Words              []string `json:"words"`
}

type Command string

// Commands
const (
    ADD_WORD Command = "add_word"
    STOP     Command = "stop"
)

func GenerateWords(config *config.Model, cmd chan Command) *[]string {
    language := fetchLanguage(config.Language)

    n := len(language.Words)

    var nextWord func() string
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

        nextWord = func() string {
            r := rand.Intn(totalWeight)
            // Use binary search to find the index of the chosen word
            idx := sort.Search(n, func(i int) bool { return cumulativeWeights[i] > r })
            return language.Words[idx]
        }
    } else {
        nextWord = func() string {
            return language.Words[rand.Intn(n)]
        }
    }

	switch config.Mode {
	case "time":
        return language.generateTimeTest(nextWord, cmd)
    case "words":
        return language.generateWordTest(nextWord, config.Words, cmd)
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

func (this *languageModel) generateTimeTest(nextWord func() string, cmd chan Command) *[]string {
    var generated []string
    for range INIT_WORDS {
        generated = append(generated, nextWord())
    }

    go commandHandler(cmd, &generated, nextWord)
    return &generated
}

func (this *languageModel) generateWordTest(nextWord func() string, words int, cmd chan Command) *[]string {
    var generated []string

    wordsToGenerate := words
    if words == 0 {
        wordsToGenerate = INIT_WORDS
    }

    for range wordsToGenerate {
        generated = append(generated, nextWord())
    }

    go commandHandler(cmd, &generated, nextWord)
    return &generated
}

func commandHandler(cmd chan Command, generated *[]string, nextWord func() string) {
    for {
        switch <-cmd {
        case ADD_WORD:
            *generated = append(*generated, nextWord())
        case STOP:
            return
        }
    }
}
