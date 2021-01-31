剑指 Offer 55 - II. 平衡二叉树

地址：[剑指 Offer 55 - II. 平衡二叉树](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/)

> 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
>
>  
>
> 示例 1:
>
> 给定二叉树 [3,9,20,null,null,15,7]
>
>     3
>    / \
>   9  20
>     /  \
>    15   7
> 返回 true 。
>
> 示例 2:
>
> 给定二叉树 [1,2,2,3,3,null,null,4,4]
>
>        1
>       / \
>      2   2
>     / \
>    3   3
>   / \
>  4   4
> 返回 false 。
>
>  
>
> 限制：
>
> 1 <= 树的结点个数 <= 10000

``` 

```

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return recur(root) != -1
}

func recur(root *TreeNode) int {
    if root == nil  {return 0}
    left := recur(root.Left)
    if left == -1   {return -1}
    right := recur(root.Right)
    if right == -1  {return  -1}
    if abs(left - right) < 2 {
        return max(left, right) + 1
    } else {
        return -1
    }
}



func max (a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func abs(x int) int {
    if x < 0 {
        x = -x
    }
    return x
}


```

