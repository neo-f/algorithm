package _917_仅仅反转字母

import "unicode"

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	for left, right := 0, len(s)-1; left < right; {
		if unicode.IsLetter(runes[left]) && unicode.IsLetter(runes[right]) {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
		if !unicode.IsLetter(runes[left]) {
			left++
		}
		if !unicode.IsLetter(runes[right]) {
			right--
		}
	}
	return string(runes)
}
