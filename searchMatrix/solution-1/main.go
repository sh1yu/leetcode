package main

func main() {
    matrix1 := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
    result1 := searchMatrix(matrix1, 5)
    println(result1)

    matrix2 := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
    result2 := searchMatrix(matrix2, 20)
    println(result2)
}

// 思路：从左下角开始，如果当前值比搜索值大则向上查找；如果当前值比搜索值小则向右查找。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix) - 1
    col := 0
    for ; row >= 0 && col < len(matrix[0]); {
        if matrix[row][col] > target {
            row--
            continue
        }
        if matrix[row][col] < target {
            col++
            continue
        }
        return true
    }
    return false
}
