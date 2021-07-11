package main

import "fmt"

func main() {
    nums1 := []int{1, 2, 3, 0, 0, 0}
    m1 := 3
    nums2 := []int{2, 5, 6}
    m2 := 3
    merge(nums1, m1, nums2, m2)
    fmt.Println(nums1)

    nums1 = []int{0}
    nums2 = []int{1}
    merge(nums1, 0, nums2, 1)
    fmt.Println(nums1)
}

// 思路：直接从后向前双指针遍历nums1，逐一拷贝即可
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
    i := m - 1
    j := n - 1
    for cur := m + n - 1; cur >= 0; cur-- {
        if j < 0 {
            break
        }

        if i < 0 || nums1[i] < nums2[j] {
            nums1[cur] = nums2[j]
            j--
            continue
        }

        nums1[cur] = nums1[i]
        i--
    }
}
