剑指 Offer 07. 重建二叉树

地址：[剑指 Offer 07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

问题描述：

>输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
>
> 
>
>例如，给出
>
>前序遍历 preorder = [3,9,20,15,7]
>中序遍历 inorder = [9,3,15,20,7]
>返回如下的二叉树：
>
>    3
>   / \
>  9  20
>    /  \
>   15   7
>
>
>限制：
>
>0 <= 节点个数 <= 5000
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
//注意算清两侧长度
object Solution {
    def buildTree(preorder: Array[Int], inorder: Array[Int]): TreeNode = {
        val root = build(preorder, inorder, 0, preorder.length-1, 0, inorder.length-1)
        return root
    }

    def build(preorder: Array[Int], inorder: Array[Int], pl: Int, pr: Int, il: Int, ir:Int): TreeNode = {
        if (pl > pr || il > ir) return null
        val pivot = inorder.indexOf(preorder(pl))
        val root = new TreeNode(preorder(pl))
        root.left = build(preorder, inorder, pl+1, pl + (pivot - il) , il, pivot-1)
        root.right = build(preorder, inorder, pl + (pivot - il) + 1, pr, pivot+1, ir)
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {return nil}

    index := 0
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            index = i
            break
        }
    } 

    root := &TreeNode{Val: preorder[0]}
    root.Left = buildTree(preorder[1:index+1], inorder[0:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])

    return root
}
```

