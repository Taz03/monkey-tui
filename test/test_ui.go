package test

import "strings"

func (this *Test) View() string {
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

func (this *Test) partitionedRows() (rows []rowModel) {
    var row rowModel
    for i := range this.Words {
        var word string
        if i >= len(this.typedWords) || len(this.Words[i]) >= len(this.typedWords[i]) {
            word = this.Words[i]
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

func (this *Test) focusedRows(rows []rowModel) (focusedRows []rowModel) {
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
                    for i := len(rows) - 3; i >= 0 && i < len(rows); i++ {
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

func (this *Test) renderWord(i int) string {
    if i >= len(this.typedWords) {
        return this.Config.StyleUntyped(style.Copy(), this.Words[i] + " ").Render()
    }

    wordStyle := style.Copy()
    if i < this.pos[0] && this.Words[i] != this.typedWords[i] {
        this.Config.StyleWrongWordUnderline(wordStyle)
    }

    var renderedWord strings.Builder
    for j, char := range this.typedWords[i] {
        str := string(char)
        switch {
        case j >= len(this.Words[i]):
            renderedWord.WriteString(this.Config.StyleErrorExtra(wordStyle.Copy(), str).Render())
        case str != string(this.Words[i][j]):
            renderedWord.WriteString(this.Config.StyleError(wordStyle.Copy(), str, string(this.Words[i][j])).Render())
        default:
            renderedWord.WriteString(this.Config.StyleCorrect(wordStyle.Copy(), str).Render())
        }
    }

    var remainingWord string
    if len(this.Words[i]) > len(this.typedWords[i]) {
        remainingWord += this.Words[i][len(this.typedWords[i]):]
    }
    remainingWord += " "

    if i < this.pos[0] {
        renderedWord.WriteString(this.Config.StyleUntyped(wordStyle, remainingWord).Render())

        return renderedWord.String()
    }

    caret.SetChar(string(remainingWord[0]))
    renderedWord.WriteString(caret.View())

    renderedWord.WriteString(this.Config.StyleUntyped(wordStyle, remainingWord[1:]).Render())

    return renderedWord.String()
}
