剑指 Offer 26. 树的子结构

地址：[剑指 Offer 26. 树的子结构](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

问题描述：

>输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
>
>B是A的子结构， 即 A中有出现和B相同的结构和节点值。
>
>例如:
>给定的树 A:
>
>     3
>    / \
>   4   5
>  / \
> 1   2
>给定的树 B：
>
>   4 
>  /
> 1
>返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
>
>示例 1：
>
>输入：A = [1,2,3], B = [3,1]
>输出：false
>示例 2：
>
>输入：A = [3,4,5,1,2], B = [4,1]
>输出：true
>限制：
>
>0 <= 节点个数 <= 10000
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
    def isSubStructure(A: TreeNode, B: TreeNode): Boolean = {
       if (A == null || B == null) {return false}
       return isSub(A, B) || isSubStructure(A.left, B) || isSubStructure(A.right, B)
    }

    def isSub(A: TreeNode, B: TreeNode): Boolean = {
        if (B == null) {return true}
        if (A == null || A.value != B.value) {return false}
        return isSub(A.left, B.left) && isSub(A.right, B.right)
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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil {return false}
    return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A *TreeNode, B *TreeNode) bool {
    if B == nil {return true}
    if A == nil || A.Val != B.Val {return false}
    return isSub(A.Left, B.Left) && isSub(A.Right, B.Right)
}
```

