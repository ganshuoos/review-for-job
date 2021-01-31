剑指 Offer 32 - III. 从上到下打印二叉树 III

地址：[剑指 Offer 32 - III. 从上到下打印二叉树 III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)

>请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
>  [20,9],
>  [15,7]
>]
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
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {return res}

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    cnt := 1

    for len(queue) > 0 {
        floorRes := make([]int, 0)

        for i := len(queue); i > 0; i-- {
            curTreeNode := queue[0]
            queue = queue[1:]
            floorRes = append(floorRes, curTreeNode.Val)
            if curTreeNode.Left != nil {queue = append(queue, curTreeNode.Left)}
            if curTreeNode.Right != nil {queue = append(queue, curTreeNode.Right)}
        } 

        if (cnt & 1) == 0 {
            reverse(floorRes)
        }
        res = append(res, floorRes)
        cnt += 1
    }
    return res
}

func reverse(floorRes []int)  {
    for i, j := 0, len(floorRes)-1; i <= j; i, j = i+1, j-1 {
        floorRes[i], floorRes[j] = floorRes[j], floorRes[i]
    }
}
```

