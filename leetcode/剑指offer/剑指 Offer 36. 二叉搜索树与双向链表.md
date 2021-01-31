剑指 Offer 36. 二叉搜索树与双向链表

地址：[剑指 Offer 36. 二叉搜索树与双向链表](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/)

>输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
>
> 
>
>为了让您更好地理解问题，以下面的二叉搜索树为例：
>
> 
>
>
>
> 
>
>我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
>
>下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
>
> 
>
>
>
> 
>
>特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。
>

``` scala

```

```go
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	root := &TreeNode{4, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{5, nil, nil}
	node3 := &TreeNode{1, nil, nil}
	node4 := &TreeNode{3, nil, nil}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	head := treeToDoublyList(root)
	tail := head.Left
	//从头开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", head.Val)
		head = head.Right
	}
	fmt.Println("-------------------")
	//从尾开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", tail.Val)
		tail = tail.Left
	}
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {return root}
	var pre *TreeNode
	dfs(root, &pre)
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

func dfs(head *TreeNode, pre **TreeNode) {
	if head == nil {return}
	dfs(head.Left, pre)
	if (*pre) != nil {
		head.Left = *pre
		(*pre).Right = head
	}
	*pre = head
	dfs(head.Right, pre)
}

```

