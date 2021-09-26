package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {

	ans := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			if nums[j]+nums[k]+nums[i] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[j]+nums[k]+nums[i] < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return ans
}
