func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i, ch := range s {
		count[ch-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
