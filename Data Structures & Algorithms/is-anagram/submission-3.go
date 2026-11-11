/*
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

Two strings are anagrams if they contain the same characters, with each character appearing the same number of times, regardless of order.


Example 1:

Input: s = "racecar", t = "carrace"

Output: true
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
