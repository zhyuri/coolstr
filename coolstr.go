package coolstr

type formater func(string) string

func coolStr(input string) int64 {
	length := len(input)
	if length <= 0 {
		return 0
	}

	wordMap := make(map[string]int)
	var wordCount int64
	lastNan := 0
	for i, char := range input {
		// change line \n
		if char == 10 {
			println("meet new line")
			wordMap[input[lastNan:i]] = 1
		}
		// space
		if char == 32 {
			println("meet space")
		}

	}
	return wordCount
}

func formatString(token string) string {
	return ""
}
func formatNum(token string) string {
	return ""
}
