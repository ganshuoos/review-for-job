剑指 Offer 33. 二叉搜索树的后序遍历序列

地址：[剑指 Offer 33. 二叉搜索树的后序遍历序列](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/)

>输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
>
> 
>
>参考以下这颗二叉搜索树：
>
>     5
>    / \
>   2   6
>  / \
> 1   3
>示例 1：
>
>输入: [1,6,3,2,5]
>输出: false
>示例 2：
>
>输入: [1,3,2,6,5]
>输出: true
>
>
>提示：
>
>数组长度 <= 1000
>

```go
//递归法 利用二叉搜索树处理 左儿子 < 根节点 < 右儿子
//利用节点值与根节点值比较，分隔出左右子树 对左右子树进行递归
func verifyPostorder(postorder []int) bool {
    return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, left, right int) bool {
    if left >= right {return true}
    cur := left
    for postorder[cur] < postorder[right] {
        cur += 1
    }
    mid := cur
    for postorder[cur] > postorder[right] {
        cur += 1
    }
    return cur == right && recur(postorder, left, mid-1) && recur(postorder, mid, right-1)
}

//利用单调栈，模拟一个没有右子树的二叉搜索树，不断去掉访问到的右子树更新root节点，直到全部节点出入栈一次
func verifyPostorder(postorder []int) bool {
    root := math.MaxInt32
    stack := make([]int, 0)

    for i := len(postorder)-1; i >= 0; i-- {
        if postorder[i] > root {return false}
        for len(stack) > 0 && postorder[i] < stack[len(stack)-1] {
            root = stack[len(stack)-1]
            stack = stack[0:len(stack)-1]
        }
        stack = append(stack, postorder[i])
    }

    return true
}
```

