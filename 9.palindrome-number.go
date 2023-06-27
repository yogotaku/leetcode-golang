/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// reverse x
	r := 0
	for i := x; i > 0; i /= 10 {
		r = r*10 + i%10
	}

	return x == r
}

// @lc code=end
