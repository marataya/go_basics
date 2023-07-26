//315. Count of Smaller Numbers After Self
//Hard
//8.3K
//225
//Companies
//Given an integer array nums, return an integer array counts where counts[i] is the number of smaller elements to the right of nums[i].
//
//
//
//Example 1:
//
//Input: nums = [5,2,6,1]
//Output: [2,1,1,0]
//Explanation:
//To the right of 5 there are 2 smaller elements (2 and 1).
//To the right of 2 there is only 1 smaller element (1).
//To the right of 6 there is 1 smaller element (1).
//To the right of 1 there is 0 smaller element.
//Example 2:
//
//Input: nums = [-1]
//Output: [0]
//Example 3:
//
//Input: nums = [-1,-1]
//Output: [0,0]
//
//
//Constraints:
//
//1 <= nums.length <= 105
//-104 <= nums[i] <= 104

package main

func CountSmaller(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				result[i]++
			}
		}
	}
	return result
}
