package util

import "unicode"

func ToRuneSlice(s string) []rune {
	runeSlice := make([]rune, len(s))

	for i, r := range s {
		runeSlice[i] = r
	}

	return runeSlice
}

func IsWhiteSpace(s string) bool {
	isWhiteSpace := true

	for _, r := range s {
		isWhiteSpace = isWhiteSpace && unicode.IsSpace(r)
		if !isWhiteSpace {
			break
		}
	}

	return isWhiteSpace
}
