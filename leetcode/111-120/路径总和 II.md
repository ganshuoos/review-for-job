113. 路径总和 II

源地址：[113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)

问题描述：

>给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
>
>说明: 叶子节点是指没有子节点的节点。
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

``` go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var path []int
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
    res = make([][]int, 0)
    if root == nil {return res}
    path = make([]int, 0)
    dfs(root, targetSum, path)
    return res
}

func dfs(root *TreeNode, targetSum int, path []int) {
    path = append(path, root.Val)
    targetSum -= root.Val

    if root.Left != nil || root.Right != nil {
        if root.Left != nil {
            dfs(root.Left, targetSum , path)
        }

        if root.Right != nil {
            dfs(root.Right, targetSum, path)
        }
    } else {
        if targetSum == 0 {
             tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
    }
    
    path = path[:len(path)-1] 
}
```



