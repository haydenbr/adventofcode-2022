package util

func ToRuneSlice(s string) []rune {
	runeSlice := make([]rune, len(s))

	for i, r := range s {
		runeSlice[i] = r
	}

	return runeSlice
}
