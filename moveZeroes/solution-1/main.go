package main

import "fmt"

func main() {
	num := []int{0, 1, 0, 3, 12}
	moveZeroes(num)
	fmt.Println(num)
}

// 基本思想：直接双指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func moveZeroes(nums []int) {
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
