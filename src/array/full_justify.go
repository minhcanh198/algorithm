package array

// https://leetcode.com/problems/text-justification/description/?envType=study-plan-v2&envId=top-interview-150
func FullJustify(words []string, maxWidth int) []string {
	var result []string

	wordLen := 0
	fromIdx := 0
	for idx := 0; idx < len(words); {
		if wordLen == 0 {
			fromIdx = idx
		}

		wordLen += len(words[idx])
		minLen := wordLen + idx - fromIdx
		if minLen > maxWidth {
			lineText := justifyText(words, maxWidth, fromIdx, idx-1)
			result = append(result, lineText)
			wordLen = 0
			continue
		}

		if idx == len(words)-1 {
			lineText := justifyText(words, maxWidth, fromIdx, idx)
			result = append(result, lineText)
			break
		}

		if minLen == maxWidth {
			lineText := justifyText(words, maxWidth, fromIdx, idx)
			result = append(result, lineText)
			wordLen = 0
			//idx++
			//continue
		}

		idx++
	}

	return result
}

func justifyText(words []string, maxWidth int, fromIdx, toIdx int) string {
	isLastLine := toIdx == len(words)-1
	if isLastLine {
		insertText := ""
		for i := fromIdx; i <= toIdx; i++ {
			if i == fromIdx {
				insertText = words[i]
				continue
			}
			insertText += " " + words[i]
		}
		remainSpace := maxWidth - len(insertText)
		for remainSpace > 0 {
			insertText = insertText + " "
			remainSpace--
		}
		return insertText
	}

	insertText := ""

	wordLen := 0
	for i := fromIdx; i <= toIdx; i++ {
		wordLen += len(words[i])
	}

	minLen := wordLen + toIdx - fromIdx
	redundance := maxWidth - minLen
	maxSpaceAdd := 0
	if fromIdx == toIdx {
		maxSpaceAdd = redundance
		insertText = words[fromIdx]
		for j := 0; j < maxSpaceAdd; j++ {
			insertText += " "
			redundance--
		}
		return insertText
	}

	for i := fromIdx; i <= toIdx; i++ {
		if i == fromIdx {
			insertText = words[i]
			continue
		}

		if redundance%(toIdx-i+1) == 0 {
			maxSpaceAdd = redundance / (toIdx - i + 1)
		} else if redundance < (toIdx - i + 1) {
			maxSpaceAdd = 1
		} else {
			maxSpaceAdd = redundance/(toIdx-i+1) + 1
		}

		for j := 0; j < maxSpaceAdd; j++ {
			if redundance > 0 {
				insertText += " "
				redundance--
			}
		}
		insertText += " " + words[i]
	}

	return insertText
}
