package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println((countVietnameseCharacters("hwfdawhwhcoomddfgwdcw")))
}

func countVietnameseCharacters(str string) int {
	telexPatterns := []string{"aw", "aa", "dd", "ee", "oo", "ow", "w"}
	count := 0
	foundChars := []string{}

	for i := 0; i < len(str); i++ {
		if i+1 < len(str) {
			twoChar := string(str[i : i+2])
			if slices.Contains(telexPatterns, twoChar) {
				count++
				i++
				foundChars = append(foundChars, twoChar)
				continue
			}
		}

		if str[i] == 'w' {
			count++
			foundChars = append(foundChars, string(str[i]))
			continue
		}

	}

	fmt.Println(foundChars)

	return count
}
