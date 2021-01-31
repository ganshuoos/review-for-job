剑指 Offer 27. 二叉树的镜像

地址：[剑指 Offer 27. 二叉树的镜像](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

问题描述：

>请完成一个函数，输入一个二叉树，该函数输出它的镜像。
>
>例如输入：
>
>     4
>   /   \
>  2     7
> / \   / \
>1   3 6   9
>镜像输出：
>
>     4
>   /   \
>  7     2
> / \   / \
>9   6 3   1
>
> 
>
>示例 1：
>
>输入：root = [4,2,7,1,3,6,9]
>输出：[4,7,2,9,6,3,1]
>
>
>限制：
>
>0 <= 节点个数 <= 1000
>

``` scala
/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */
object Solution {
    def mirrorTree(root: TreeNode): TreeNode = {
        if (root == null) return root
        val left = mirrorTree(root.right)
        val right = mirrorTree(root.left)
        root.left = left
        root.right = right
        return root
    }
}
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
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {return root}
    root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
    return root 
}
```

