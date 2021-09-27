package main

import "fmt"

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	groupParents := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				groupParents[i*n+j] = i*n + j
				count++
			} else {
				groupParents[i*n+j] = -1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				if j < n-1 && grid[i][j+1] == '1' {
					unite(groupParents, &count, i*n+j, i*n+j+1)
				}
				if i < m-1 && grid[i+1][j] == '1' {
					unite(groupParents, &count, i*n+j, (i+1)*n+j)
				}
			}
		}
	}

	return count
}
func find(groupParents []int, i int) int {
	if i >= 0 && groupParents[i] >= 0 && groupParents[i] != i {
		groupParents[i] = find(groupParents, groupParents[i])
	}
	return groupParents[i]
}

func unite(groupParents []int, count *int, i, j int) {
	rooti := find(groupParents, i)
	rootj := find(groupParents, j)
	if rooti >= 0 && rootj >= 0 && rooti != rootj {
		groupParents[rootj] = rooti
		*count--
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}
