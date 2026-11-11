func hasDuplicate(nums []int) bool {
    mpp := make(map[int]bool)
	for _, v := range nums {
		if mpp[v] {
			return true
		}
		mpp[v] = true
	}
	return false
}
