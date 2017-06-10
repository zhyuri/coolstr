package coolstr

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"regexp"
)

// CoolStr get a string as input and ouput the word count of the string
func CoolStr(input string) int {
	if len(input) <= 0 {
		return 0
	}
	input += " \n"

	var (
		wordCount  int
		needConcat bool // need concatenate through two lines or not
		lastNan    int  // last ' ' and '\n' char
		word       string
	)
	wordMap := make(map[string]int)
	for i, char := range input {
		if char != '\n' && char != ' ' {
			continue
		}
		if word += strings.TrimSpace(input[lastNan:i]); word == "" {
			continue
		}
		lastNan = i + 1
		if char == '\n' && strings.HasSuffix(word, "-") {
			// trim suffix "-"
			word = word[:len(word)-len("-")]
			lastRune, _ := utf8.DecodeLastRuneInString(word)
			needConcat = !isNotAlphabet(lastRune)
		}
		if needConcat {
			needConcat = false
			continue
		}
		// trim suffix punctuation
		firstRune, _ := utf8.DecodeRuneInString(word[:1])
		if !isCapital(firstRune) {
			word = strings.TrimRightFunc(word, isNotAlphabet)
		}
		fmt.Printf("|%d\t%d\t%s\t|\n", lastNan, i, word)
		wordCount += process(wordMap, !isAlphabet(firstRune), word)
		word = ""
	}
	fmt.Printf("-------------------\t|\n")
	return wordCount
}

var notAlphaRegexp = regexp.MustCompile("[^a-zA-Z0-9.-]")

func process(wordMap map[string]int, isNum bool, token string) (count int) {
	if isNum {
		token = strings.Replace(token, "-", "", -1)
	}
	if token = notAlphaRegexp.ReplaceAllLiteralString(token, ""); token == "" {
		return
	}
	token = strings.ToLower(token)
	wordMap[token]++
	if wordMap[token] <= 1 {
		count = 1
	}
	return
}
func isCapital(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
func isAlphabet(r rune) bool {
	return (r >= 'a' && r <= 'z') || isCapital(r)
}
func isNotAlphabet(r rune) bool {
	return !(isAlphabet(r) || (r >= '0' && r <= '9'))
}
