package coolstr

import (
	"strings"
	"unicode/utf8"
)

// CoolStr get a string as input and ouput the word count of the string
func CoolStr(input string) int {
	length := len(input)
	if length <= 0 {
		return 0
	}

	var (
		wordCount int
		needCat   bool // need concatenate through two lines or not
		lastNan   int  // last ' ' and '\n' char
	)
	wordMap := make(map[string]int)
	for i, char := range input {
		// space
		if char != '\n' && char != ' ' && i-lastNan <= 1 {
			continue
		}
		word := input[lastNan:i]
		// newline \n
		if char == '\n' {
			word, needCat = trimReturn(word)
		}
		if needCat {
			needCat = false
			continue
		}
		lastNan = i
		wordCount += process(wordMap, word)

	}
	return wordCount
}

type formater func(string) string

var formatPipline = []formater{
	formatNum,
	formatString,
}

func process(wordMap map[string]int, token string) int {
	for _, formatFunc := range formatPipline {
		token = formatFunc(token)
	}
	wordMap[token]++
	if token != "" && wordMap[token] <= 1 {
		return 1
	}
	return 0
}

func trimReturn(token string) (newToken string, needCat bool) {
	newToken = strings.TrimSpace(token)
	if strings.HasSuffix(token, "-") {
		newToken = strings.TrimSuffix(newToken, "-")
		lastRune, _ := utf8.DecodeLastRuneInString(token)
		needCat = (lastRune >= 'a' && lastRune <= 'z') || (lastRune >= 'A' && lastRune <= 'Z')
	}
	return
}

func formatString(token string) string {
	return token
}
func formatNum(token string) string {
	return token
}
