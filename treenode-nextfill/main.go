package main

import "fmt"

// Definition for a binary tree node.
type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}

func main() {
    node7 := &Node{
        Val:   7,
        Left:  nil,
        Right: nil,
    }
    node4 := &Node{
        Val:   4,
        Left:  node7,
        Right: nil,
    }
    node5 := &Node{
        Val:   5,
        Left:  nil,
        Right: nil,
    }
    node2 := &Node{
        Val:   2,
        Left:  node4,
        Right: node5,
    }

    node8 := &Node{
        Val:   8,
        Left:  nil,
        Right: nil,
    }
    node6 := &Node{
        Val:   6,
        Left:  nil,
        Right: node8,
    }
    node3 := &Node{
        Val:   3,
        Left:  nil,
        Right: node6,
    }
    node1 := &Node{
        Val:   1,
        Left:  node2,
        Right: node3,
    }

    connect(node1)
    fmt.Println(node7.Next)
}

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    dfs(root, nil)
    return root
}

//错误的算法。只适用于完全二叉树
func dfs(node *Node, next *Node) {
    if node == nil {
        return
    }
    node.Next = next

    realNext := next
    var realNextChild *Node
    for {
        if realNext == nil {
            break
        }
        if realNext.Left != nil {
            realNextChild = realNext.Left
            break
        }
        if realNext.Right != nil {
            realNextChild = realNext.Right
            break
        }
        realNext = realNext.Next
    }

    //fmt.Println(realNextChild)

    if node.Left != nil {
        if node.Right != nil {
            dfs(node.Left, node.Right)
            dfs(node.Right, realNextChild)
        } else {
            dfs(node.Left, realNextChild)
        }
    } else {
        if node.Right != nil {
            dfs(node.Right, realNextChild)
        }
    }
}
