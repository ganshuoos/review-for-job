25. K 个一组翻转链表

源地址：[25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

问题描述：

>给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
>
>k 是一个正整数，它的值小于或等于链表的长度。
>
>如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
>
> 
>
>示例：
>
>给你这个链表：1->2->3->4->5
>
>当 k = 2 时，应当返回: 2->1->4->3->5
>
>当 k = 3 时，应当返回: 3->2->1->4->5
>
> 
>
>说明：
>
>你的算法只能使用常数的额外空间。
>你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummyHead := &ListNode{0, head}
    for p := dummyHead; ; {
        q := p
        for i := 0; i < k && q != nil; i++ {q = q.Next}
        if q == nil {break}

        a := p.Next
        b := a.Next
        for i:=0; i < k-1; i++ {
            c := b.Next
            b.Next = a
            a = b 
            b = c
        }

        c := p.Next
        p.Next = a
        c.Next = b 
        p = c
    }
    return  dummyHead.Next
}
```



