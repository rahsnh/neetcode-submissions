/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true
*/
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
