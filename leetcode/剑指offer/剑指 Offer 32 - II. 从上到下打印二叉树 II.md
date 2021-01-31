剑指 Offer 32 - II. 从上到下打印二叉树 II

地址：[剑指 Offer 32 - II. 从上到下打印二叉树 II](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/)

>从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
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
>返回其层次遍历结果：
>
>[
>  [3],
>  [9,20],
>  [15,7]
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
func levelOrder(root *TreeNode) [][]int {
    //floorRes := make([]int, 0)
    res := make([][]int, 0)
    if root == nil {return res}

    queue := make([]*TreeNode, 0)
    queue = append(queue, root) 
    for len(queue) > 0 {
        floorRes := make([]int, 0)      
        //获取当前队列中剩余元素个数
        for i := len(queue); i > 0; i-- {
            curTreeNode := queue[0]
            queue = queue[1:]
            floorRes = append(floorRes, curTreeNode.Val)
            if curTreeNode.Left != nil {queue = append(queue, curTreeNode.Left)}
            if curTreeNode.Right != nil {queue = append(queue, curTreeNode.Right)}
        }
        res = append(res, floorRes)
    }
    return res
}
```

