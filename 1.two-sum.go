
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
// @lc code=start

// Your runtime beats 56.59 % of golang submissions
// Your memory usage beats 12.82 % of golang submissions (4.6 MB)
func twoSum(nums []int, target int) []int {
	ni := make(map[int]int, len(nums))
	output := make([]int, 2)
	for i, n := range nums {
		ni[n] = i
	}

	for i, n := range nums {
		t := target - n
		oi, ok := ni[t]
		if ok && i != oi {
			output[0] = i
			output[1] = oi
			break
		}
	}

	return output
}

// @lc code=end
