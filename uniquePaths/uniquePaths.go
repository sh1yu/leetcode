package main

import "fmt"

var dp [][]int

func uniquePaths(m int, n int) int {
	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	//for i := 1; i < m; i++ {
	//    for j := 1; j < n; j++ {
	//        dp[i][j] = dp[i-1][j] + dp[i][j-1]
	//    }
	//}

	dfs1(m, n, m-1, n-1)

	return dp[m-1][n-1]
}

func dfs1(m, n, i, j int) {
	if m == 1 || n == 1 {
		return
	}

	if dp[i-1][j] == 0 {
		dfs1(m, n, i-1, j)
	}
	if dp[i][j-1] == 0 {
		dfs1(m, n, i, j-1)
	}

	dp[i][j] = dp[i-1][j] + dp[i][j-1]
}

func main() {
	fmt.Println(uniquePaths(1, 2))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(51, 9))
}
