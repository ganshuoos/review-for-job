剑指 Offer 37. 序列化二叉树

地址：[剑指 Offer 37. 序列化二叉树](https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/)

>请实现两个函数，分别用来序列化和反序列化二叉树。
>
>示例: 
>
>你可以将以下二叉树：
>
>    1
>   / \
>  2   3
>     / \
>    4   5
>
>序列化为 "[1,2,3,null,null,4,5]"
>注意：本题与主站 297 题相同：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

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
import (
    "strconv"
    "fmt"
)

var path string
//var u int

type Codec struct{
    list []string
}

func Constructor() Codec {
    return Codec{}    
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    dfsSer(root)
    return path    
}

func dfsSer(root *TreeNode) {
    if root == nil {
        path += "#,"
    } else {
        path += strconv.Itoa(root.Val) + ","      
        dfsSer(root.Left)
        dfsSer(root.Right)
    }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode { 
    u := 0
    return dfsDse(data, &u)
}

func dfsDse(data string, u *int) *TreeNode {
    if data[*u] == '#' {
        *u += 2
        return nil
    } else {
        k := *u
        for data[*u] != ',' {*u += 1}
        root := &TreeNode{}
        root.Val, _ = strconv.Atoi(data[k:*u])
        *u += 1
        root.Left = dfsDse(data, u)
        root.Right = dfsDse(data, u)
        return root 
    }

}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
```

