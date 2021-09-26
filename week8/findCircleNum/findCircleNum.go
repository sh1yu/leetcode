package main

import "fmt"

func findCircleNum(isConnected [][]int) int {

	count := len(isConnected)
	parents := make([]int, len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		parents[i] = i
	}

	for i := 0; i < len(isConnected)-1; i++ {
		for j := i + 1; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				parent_i := i
				for parent_i != parents[parent_i] {
					next_i := parents[parent_i]
					parents[parent_i] = parents[next_i]
					parent_i = next_i
				}
				parent_j := j
				for parent_j != parents[parent_j] {
					next_j := parents[parent_j]
					parents[parent_j] = parents[next_j]
					parent_j = next_j
				}
				if parent_i == parent_j {
					continue
				}
				parents[parent_i] = parent_j
				count--
			}
		}
	}

	return count
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
