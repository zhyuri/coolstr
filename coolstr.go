package coolstr

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// CoolStr get a string as input and ouput the word count of the string
func CoolStr(input string) int {
	length := len(input)
	if length <= 0 {
		return 0
	}
	input += " \n"

	var (
		wordCount  int
		needConcat bool // need concatenate through two lines or not
		lastNan    = -1 // last ' ' and '\n' char
	)
	wordMap := make(map[string]int)
	for i, char := range input {
		if char != '\n' && char != ' ' {
			continue
		}
		word := strings.TrimSpace(input[lastNan+1 : i])
		if char == '\n' {
			word, needConcat = trimReturn(word)
		}
		if needConcat {
			needConcat = false
			continue
		}
		lastNan = i
		fmt.Printf("|%d\t%d\tword: %s|\n", lastNan+1, i, word)
		wordCount += process(wordMap, word)

	}
	fmt.Println("----------------")
	return wordCount
}

type valve func(string) string

var pipline = []valve{
	valveString,
	valveNum,
}

func process(wordMap map[string]int, token string) (count int) {
	for _, valveFunc := range pipline {
		if token = valveFunc(token); token == "" {
			return
		}
	}
	wordMap[token]++
	if wordMap[token] <= 1 {
		count = 1
	}
	return
}

func trimReturn(token string) (output string, needCat bool) {
	output = token
	if !strings.HasSuffix(output, "-") {
		return
	}
	output = output[:len(output)-len("-")]
	lastRune, _ := utf8.DecodeLastRuneInString(output)
	needCat = (lastRune >= 'a' && lastRune <= 'z') || (lastRune >= 'A' && lastRune <= 'Z')
	return
}

func valveString(token string) string {
	return token
}

func valveNum(token string) string {
	return token
}
