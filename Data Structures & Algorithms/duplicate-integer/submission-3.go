func hasDuplicate(nums []int) bool {
    duplicateMap := make(map[int]bool)
	for _, v := range nums {
		if exists := duplicateMap[v]; exists {
			return true
		}
		duplicateMap[v] = true
	}
	return false
}
