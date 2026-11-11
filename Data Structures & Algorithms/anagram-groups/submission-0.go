/*
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
*/
func groupAnagrams(strs []string) [][]string {
	anagramsGrp := make(map[[26]int][]string)

	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[c-'a']++
		}
		anagramsGrp[count] = append(anagramsGrp[count], s)
	}

	var result [][]string
	for _, v := range anagramsGrp {
		result = append(result, v)
	}

	return result
}
