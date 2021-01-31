剑指 Offer 32 - I. 从上到下打印二叉树

地址：[剑指 Offer 32 - I. 从上到下打印二叉树](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/)

>从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
>
> 
>
>例如:
>给定二叉树: [3,9,20,null,null,15,7],
>
>    3
>   / \
>  9  20
>    /  \
>   15   7
>返回：
>
>[3,9,20,15,7]
>
>
>提示：
>
>节点总数 <= 1000
>

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
func levelOrder(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        res = append(res, queue[0].Val)
        if queue[0].Left != nil {
            queue = append(queue, queue[0].Left)
        }
        if queue[0].Right != nil {
            queue = append(queue, queue[0].Right)
        }
        queue = queue[1:]
    }
    return res
}
```

