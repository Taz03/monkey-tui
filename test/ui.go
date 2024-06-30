package test

import "strings"

func (m *Model) View() string {
	rows := m.focusedRows(m.partitionedRows())

	var renderedRows []string
	for _, row := range rows {
		renderedRow := strings.Builder{}
		for _, index := range row.indices {
			renderedRow.WriteString(m.renderWord(index))
		}

		if m.Width > row.length {
			renderedRow.WriteString(strings.Repeat(space, m.Width-row.length))
		}
		renderedRows = append(renderedRows, renderedRow.String())
	}

	return strings.Join(renderedRows, "\n")
}

type rowModel struct {
	length  int
	indices []int
}

func (m *Model) partitionedRows() (rows []rowModel) {
	words := *m.words

	var row rowModel
	for i := range words {
		var word string
		if i >= len(m.typedWords) || len(words[i]) >= len(m.typedWords[i]) {
			word = words[i]
		} else {
			word = m.typedWords[i]
		}

		wordLen := len(word) + 1
		if row.length+wordLen <= m.Width {
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

func (m *Model) focusedRows(rows []rowModel) (focusedRows []rowModel) {
	for i, row := range rows {
		for _, index := range row.indices {
			if m.pos[0] == index {
				if i == 0 {
					for i := 0; len(rows) > i && i < 3; i++ {
						focusedRows = append(focusedRows, rows[i])
					}
					return
				}

				if i == len(rows)-1 {
					for i := max(len(rows)-3, 0); i >= 0 && i < len(rows); i++ {
						focusedRows = append(focusedRows, rows[i])
					}
					return
				}

				return []rowModel{rows[i-1], row, rows[i+1]}
			}
		}
	}

	return
}

func (m *Model) renderWord(i int) string {
	words := *m.words

	if i >= len(m.typedWords) {
		return m.config.StyleUntyped(style, words[i]+" ").Render()
	}

	wordStyle := style
	if i < m.pos[0] && words[i] != m.typedWords[i] {
		wordStyle = m.config.StyleWrongWordUnderline(wordStyle)
	}

	var renderedWord strings.Builder
	for j, char := range m.typedWords[i] {
		str := string(char)
		switch {
		case j >= len(words[i]):
			renderedWord.WriteString(m.config.StyleErrorExtra(wordStyle, str).Render())
		case str != string(words[i][j]):
			renderedWord.WriteString(m.config.StyleError(wordStyle, str, string(words[i][j])).Render())
		default:
			renderedWord.WriteString(m.config.StyleCorrect(wordStyle, str).Render())
		}
	}

	var remainingWord string
	if len(words[i]) > len(m.typedWords[i]) {
		remainingWord += words[i][len(m.typedWords[i]):]
	}
	remainingWord += " "

	if i < m.pos[0] {
		renderedWord.WriteString(m.config.StyleUntyped(wordStyle, remainingWord).Render())

		return renderedWord.String()
	}

	caret.SetChar(string(remainingWord[0]))
	renderedWord.WriteString(caret.View())

	renderedWord.WriteString(m.config.StyleUntyped(wordStyle, remainingWord[1:]).Render())

	return renderedWord.String()
}
