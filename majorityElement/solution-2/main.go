package main

func main() {
	arr1 := []int{3, 2, 3}
	result1 := majorityElement(arr1)
	println(result1)

	arr2 := []int{2, 2, 1, 1, 1, 2, 2}
	result2 := majorityElement(arr2)
	println(result2)
}

// 思路：摩尔投票法：对拼消耗。需要找的数是大于n/2的，所以会最终胜利。
// 从0开始，遇见相同的数字加1，遇见不同的数字-1。如果变成0，重新选择下一个数字开始。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func majorityElement(nums []int) int {
	result := nums[0]
	count := 0
	for index, i := range nums {
		if result == i {
			count++
			continue
		}
		count--
		if count == 0 {
			result = nums[index+1]
		}
	}
	return result
}
