package test

import "strings"

func (this *Model) View() string {
    rows := this.focusedRows(this.partitionedRows())

    var renderedRows []string
    for _, row := range rows {
        renderedRow := strings.Builder{}
        for _, index := range row.indices {
            renderedRow.WriteString(this.renderWord(index))
        }

        if this.Width > row.length {
            renderedRow.WriteString(strings.Repeat(space, this.Width - row.length))
        }
        renderedRows = append(renderedRows, renderedRow.String())
    }

    return strings.Join(renderedRows, "\n")
}

type rowModel struct {
    length  int
    indices []int
}

func (this *Model) partitionedRows() (rows []rowModel) {
    words := *this.words

    var row rowModel
    for i := range words {
        var word string
        if i >= len(this.typedWords) || len(words[i]) >= len(this.typedWords[i]) {
            word = words[i]
        } else {
            word = this.typedWords[i]
        }
        
        wordLen := len(word) + 1
        if row.length + wordLen <= this.Width {
            row.length += wordLen
            row.indices = append(row.indices, i)
        } else {
            rows = append(rows, row)
            row = rowModel{
                length:  wordLen,
                indices: []int{i},
            }
        }
    }
    if row.length > 0 {
        rows = append(rows, row)
    }

    return
}

func (this *Model) focusedRows(rows []rowModel) (focusedRows []rowModel) {
    for i, row := range rows {
        for _, index := range row.indices {
            if this.pos[0] == index {
                if i == 0 {
                    for i := 0; len(rows) > i && i < 3; i++ {
                        focusedRows = append(focusedRows, rows[i])
                    }
                    return
                }

                if i == len(rows) - 1 {
                    for i := max(len(rows) - 3, 0); i >= 0 && i < len(rows); i++ {
                        focusedRows = append(focusedRows, rows[i])
                    }
                    return
                }

                return []rowModel{rows[i - 1], row, rows[i + 1]}
            }
        }
    }

    return
}

func (this *Model) renderWord(i int) string {
    words := *this.words

    if i >= len(this.typedWords) {
        return this.config.StyleUntyped(style, words[i] + " ").Render()
    }

    wordStyle := style
    if i < this.pos[0] && words[i] != this.typedWords[i] {
        wordStyle = this.config.StyleWrongWordUnderline(wordStyle)
    }

    var renderedWord strings.Builder
    for j, char := range this.typedWords[i] {
        str := string(char)
        switch {
        case j >= len(words[i]):
            renderedWord.WriteString(this.config.StyleErrorExtra(wordStyle, str).Render())
        case str != string(words[i][j]):
            renderedWord.WriteString(this.config.StyleError(wordStyle, str, string(words[i][j])).Render())
        default:
            renderedWord.WriteString(this.config.StyleCorrect(wordStyle, str).Render())
        }
    }

    var remainingWord string
    if len(words[i]) > len(this.typedWords[i]) {
        remainingWord += words[i][len(this.typedWords[i]):]
    }
    remainingWord += " "

    if i < this.pos[0] {
        renderedWord.WriteString(this.config.StyleUntyped(wordStyle, remainingWord).Render())

        return renderedWord.String()
    }

    caret.SetChar(string(remainingWord[0]))
    renderedWord.WriteString(caret.View())

    renderedWord.WriteString(this.config.StyleUntyped(wordStyle, remainingWord[1:]).Render())

    return renderedWord.String()
}
