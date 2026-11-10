package utils

import "unicode"

func MixCase(value string) string {
	runes := []rune(value)
	uppercaseNext := true
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			continue
		}
		if uppercaseNext {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
		uppercaseNext = !uppercaseNext
	}
	return string(runes)
}
