package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		l := intervals[i][0]
		r := intervals[i][1]
		if len(merged) == 0 || merged[len(merged)-1][1] < l {
			merged = append(merged, []int{l, r})
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], r)
		}
	}
	return merged
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
