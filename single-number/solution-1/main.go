package main

func main() {
	arr1 := []int{2, 2, 1}
	result1 := singleNumber(arr1)
	println(result1)

	arr2 := []int{4, 1, 2, 1, 2}
	result2 := singleNumber(arr2)
	println(result2)
}

// 要求：时间复杂度为线性复杂度，空间复杂度为常量
// 思路：二进制思路：所有数字依次异或，出现2次的相互抵消，最后得到的就是结果。
func singleNumber(nums []int) int {
	var result int
	for _, i := range nums {
		result ^= i
	}
	return result
}
