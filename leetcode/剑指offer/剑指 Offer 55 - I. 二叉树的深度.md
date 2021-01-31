剑指 Offer 55 - I. 二叉树的深度

地址：[剑指 Offer 55 - I. 二叉树的深度](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/)

> 输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
>
> 例如：
>
> 给定二叉树 [3,9,20,null,null,15,7]，
>
>     3
>    / \
>   9  20
>     /  \
>    15   7
> 返回它的最大深度 3 。
>
>  
>
> 提示：
>
> 节点总数 <= 10000
>

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
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

////////////////////////////
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    res := 0
    dfs(root, 1, &res)
    return res
}

func dfs(root *TreeNode, length int, res *int) {
    if root.Left == nil && root.Right == nil {
        *res = max(*res, length)
        return
    }

    if root.Left != nil {
        dfs(root.Left, length+1, res)
    }
    if root.Right != nil {
        dfs(root.Right, length+1, res)
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```

