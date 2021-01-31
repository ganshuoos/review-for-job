剑指 Offer 28. 对称的二叉树

地址：[剑指 Offer 28. 对称的二叉树](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/)

>请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
>
>例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
>
>    1
>   / \
>  2   2
> / \ / \
>3  4 4  3
>但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
>
>    1
>   / \
>  2   2
>   \   \
>   3    3
>
> 
>
>示例 1：
>
>输入：root = [1,2,2,3,4,4,3]
>输出：true
>示例 2：
>
>输入：root = [1,2,2,null,3,null,3]
>输出：false
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
    def isSymmetric(root: TreeNode): Boolean = {
        if (root == null) {return true}
        return recur(root.left, root.right)
    }

    def recur(A: TreeNode, B: TreeNode): Boolean = {
        if (A == null && B == null) {return true}
        if (A == null || B == null || A.value != B.value) {return false}
        return recur(A.left, B.right) && recur(A.right, B.left)
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
func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {return true}
    if left == nil || right == nil || left.Val != right.Val {return false}
    return recur(left.Left, right.Right) && recur(left.Right, right.Left)
}
```

