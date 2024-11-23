package anagram

func IsAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen < 1 || sLen > 50000 && tLen < 1 || tLen > 50000 {
		return false
	}

	set := map[rune]int{}

	for _, c := range s {
		current := set[c]
		if current != 0 {
			set[c] = set[c] + 1
		} else {
			set[c] = 1
		}
	}

	for _, c := range t {
		current := set[c]
		if current != 0 {
			set[c] = set[c] - 1
		} else {
			set[c] = -1
		}
	}

	for _, value := range set {
		if value != 0 {
			return false
		}
	}

	return true
}
