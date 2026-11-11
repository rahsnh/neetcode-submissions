/*
Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

Example 1:

Input: 
nums = [3,4,5,6], target = 7

Output: [0,1]
*/
func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if j, exists := prevMap[diff]; exists {
			return []int{j, i}
		}
		prevMap[v] = i
	}

	return []int{}
}
