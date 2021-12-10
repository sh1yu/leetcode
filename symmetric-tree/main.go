package main

import "fmt"

func main() {
    node3 := &TreeNode{
        Val:   3,
        Left:  nil,
        Right: nil,
    }
    node21 := &TreeNode{
        Val:   2,
        Left:  nil,
        Right: node3,
    }
    node22 := &TreeNode{
        Val:   2,
        Left:  node3,
        Right: nil,
    }
    node1 := &TreeNode{
        Val:   1,
        Left:  node21,
        Right: node22,
    }

    fmt.Println(isSymmetric(node1))
}

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return symmetric([]*TreeNode{root.Left}, []*TreeNode{root.Right})

}

func symmetric(leftChilds []*TreeNode, rightChilds []*TreeNode) bool {
    if len(leftChilds) != len(rightChilds) {
        return false
    }
    length := len(leftChilds)
    newLeft := make([]*TreeNode, 0)
    newRight := make([]*TreeNode, 0)
    for i := 0; i < length; i++ {
        if leftChilds[i] != nil {
            newLeft = append(newLeft, leftChilds[i].Left, leftChilds[i].Right)
        }
        if rightChilds[i] != nil {
            newRight = append(newRight, rightChilds[i].Left, rightChilds[i].Right)
        }

        if leftChilds[i] == nil && rightChilds[length-i-1] == nil {
            continue
        }
        if leftChilds[i] == nil || rightChilds[length-i-1] == nil {
            return false
        }
        if leftChilds[i].Val != rightChilds[length-i-1].Val {
            return false
        }
    }
    if len(newLeft) == 0 && len(newRight) == 0 {
        return true
    }
    return symmetric(newLeft, newRight)
}

func printNodes(nodes []*TreeNode) {
    for _, node := range nodes {
        if node == nil {
            fmt.Printf("<nil> ")
        } else {
            fmt.Printf("%v ", node.Val)
        }
    }
    fmt.Println()
}
