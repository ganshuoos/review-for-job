剑指 Offer 54. 二叉搜索树的第k大节点

地址：[剑指 Offer 54. 二叉搜索树的第k大节点](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)

> 给定一棵二叉搜索树，请找出其中第k大的节点。
>
>  
>
> 示例 1:
>
> 输入: root = [3,1,4,null,2], k = 1
>    3
>   / \
>  1   4
>   \
>    2
> 输出: 4
> 示例 2:
>
> 输入: root = [5,3,6,2,4,null,null,1], k = 3
>        5
>       / \
>      3   6
>     / \
>    2   4
>   /
>  1
> 输出: 4
>
>
> 限制：
>
> 1 ≤ k ≤ 二叉搜索树元素个数

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
import "fmt"
func kthLargest(root *TreeNode, k int) int {
    res := 0
    dfs(root, &k, &res)
    return res
}

func dfs(root *TreeNode, k *int, res *int) {
    if root == nil {return}
    dfs(root.Right, k, res)
    if *k == 0 {return}
    //fmt.Printf("k: %d, res: %d, root: %d\n", *k, *res, root.Val)
    *k -= 1
    if *k == 0 {*res = root.Val}
    dfs(root.Left, k, res)
}
```

