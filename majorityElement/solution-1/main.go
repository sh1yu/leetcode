package main

func main() {
	arr1 := []int{3, 2, 3}
	result1 := majorityElement(arr1)
	println(result1)

	arr2 := []int{2, 2, 1, 1, 1, 2, 2}
	result2 := majorityElement(arr2)
	println(result2)
}

// 思路：使用hashmap存储次数。如果有存在数字次数大于n/2时返回该数字。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	for _, i := range nums {
		if c, ok := countMap[i]; ok {
			countMap[i] = c + 1
		} else {
			countMap[i] = 1
		}
		if countMap[i] > len(nums)/2 {
			return i
		}
	}
	return 0
}
