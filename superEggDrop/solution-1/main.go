package main

import "math"

func main() {
    println(superEggDrop(1, 1))
    println(superEggDrop(1, 10))
    println(superEggDrop(10, 10))
    println(superEggDrop(2, 6))
    println(superEggDrop(3, 14))
}

// 思路：动态规划 + 二分查找
// 时间复杂度： O(nklogn)
func superEggDrop(k, n int) int {
    return dp(k, n)
}

var memo = make(map[int]int)

func dp(k, n int) int {
    pointer := n*100 + k
    if val, ok := memo[pointer]; ok {
        return val
    } else {
        ans := 0
        if n == 0 {
            ans = 0
        } else if k == 1 {
            ans = n
        } else {
            lo := 1
            hi := n
            for lo+1 < hi {
                x := (lo + hi) / 2
                t1 := dp(k-1, x-1)
                t2 := dp(k, n-x)
                if t1 < t2 {
                    lo = x
                } else if t1 > t2 {
                    hi = x
                } else {
                    lo = x
                    hi = x
                }
            }
            ans = 1 + int(math.Min(
                math.Max(float64(dp(k-1, lo-1)), float64(dp(k, n-lo))),
                math.Max(float64(dp(k-1, hi-1)), float64(dp(k, n-hi)))))
            memo[pointer] = ans
        }
        return ans
    }
}
