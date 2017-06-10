package coolstr

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"regexp"
)

// CoolStr get a string as input and ouput the word count of the string
func CoolStr(input string) (wordCount int) {
	if len(input) <= 0 {
		return 0
	}
	input += " \n"

	var (
		start int // last ' ' and '\n' char
		word  string
	)
	wordMap := make(map[string]int)
	for end, char := range input {
		if char != '\n' && char != ' ' {
			continue
		}
		if word += strings.TrimSpace(input[start:end]); word == "" {
			continue
		}
		start = end + 1
		if char == '\n' && strings.HasSuffix(word, "-") {
			// trim suffix "-"
			word = word[:len(word)-len("-")]
			lastRune, _ := utf8.DecodeLastRuneInString(word)
			if !isNotAlphabet(lastRune) {
				continue
			}
		}
		firstRune, _ := utf8.DecodeRuneInString(word[:1])
		// except abbreviation
		if isLowerCase(firstRune) {
			// trim suffix punctuation
			word = strings.TrimRightFunc(word, isNotAlphabet)
		}
		fmt.Printf("|%d\t%d\t%s\t|\n", start, end, word)
		wordCount += process(wordMap, !isAlphabet(firstRune), word)
		word = ""
	}
	fmt.Printf("-------------------\t|\n")
	return wordCount
}

var cleanExceptRegexp = regexp.MustCompile("[^a-zA-Z0-9.-]")

func process(wordMap map[string]int, isNum bool, token string) (count int) {
	if isNum {
		token = strings.Replace(token, "-", "", -1)
	}
	if token = cleanExceptRegexp.ReplaceAllLiteralString(token, ""); token == "" {
		return
	}
	token = strings.ToLower(token)
	wordMap[token]++
	if wordMap[token] <= 1 {
		count = 1
	}
	return
}
func isLowerCase(r rune) bool {
	return r >= 'a' && r <= 'z'
}
func isAlphabet(r rune) bool {
	return isLowerCase(r) || (r >= 'A' && r <= 'Z')
}
func isNotAlphabet(r rune) bool {
	return !(isAlphabet(r) || (r >= '0' && r <= '9'))
}
