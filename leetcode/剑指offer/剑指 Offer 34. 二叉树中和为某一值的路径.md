剑指 Offer 34. 二叉树中和为某一值的路径

地址：[剑指 Offer 34. 二叉树中和为某一值的路径](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/)

>输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
>
> 
>
>示例:
>给定如下二叉树，以及目标和 sum = 22，
>
>              5
>             / \
>            4   8
>           /   / \
>          11  13  4
>         /  \    / \
>        7    2  5   1
>返回:
>
>[
>   [5,4,11,2],
>   [5,8,4,5]
>]

``` scala

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
func pathSum(root *TreeNode, sum int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    if (root != nil) {
        dfs(root, sum, path, &res)
    }
    return res 
}

func dfs(root *TreeNode, sum int, path []int, res *[][]int) {
    path = append(path, root.Val)
    sum -= root.Val
    if root.Left == nil && root.Right == nil {
        if sum == 0 {
            validPath := make([]int, len(path))
            //注意对slice（引用类型）进行值copy
            copy(validPath, path)
            (*res) = append((*res), validPath)
        }
    } else {
        if root.Left != nil {dfs(root.Left, sum, path, res)}
        if root.Right != nil {dfs(root.Right, sum, path, res)}
    }
    path = path[0: len(path)-1]
}
```

