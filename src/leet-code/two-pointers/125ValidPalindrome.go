package two_pointers

func IsPalindrome(s string) bool {
	s = toLowerAndRemoveNonAlphabet(s)

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func toLowerAndRemoveNonAlphabet(str string) string {
	result := ""
	for _, char := range str {
		if isAlphaNumeric(char) {
			if isUpperCase(char) {
				char += 'a' - 'A'
			}
			result += string(char)
		}
	}
	return result
}

func isAlphaNumeric(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >='0' && char <='9')
}

func isUpperCase(char rune) bool {
	return char >= 'A' && char <= 'Z'
}
