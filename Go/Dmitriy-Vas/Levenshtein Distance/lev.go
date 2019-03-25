package Levenshtein_Distance

func Lev(s string, w string) int {
	sRunes := []rune(s)
	wRunes := []rune(w)

	sLength := len(s)
	wLength := len(w)

	table := make([]int, sLength + 1)

	for i := 1; i <= sLength; i++ {
		table[i] = i
	}

	for i := 1; i <= wLength; i++ {
		d := i - 1
		for j := 1; j <= sLength; j++ {
			o := table[j]
			cost := 0
			if sRunes[j - 1] != wRunes[i - 1] {
				cost = 1
			}
			table[j] = min(table[j]+1,table[j-1]+1,d+cost)
			d = o
		}
	}
	return table[sLength]
}

func min(a int, b int, c int) int {
	if c < a {
		if c < b {
			return c
		}
	} else if a < b {
		return a
	}
	return b
}
