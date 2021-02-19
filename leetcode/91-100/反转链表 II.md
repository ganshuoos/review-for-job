92. 反转链表 II

源地址：[92. 反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)

问题描述：

>反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
>
>说明:
>1 ≤ m ≤ n ≤ 链表长度。
>
>示例:
>
>输入: 1->2->3->4->5->NULL, m = 2, n = 4
>输出: 1->4->3->2->5->NULL

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummyHead := &ListNode{-1, head}
    a := dummyHead

    for i := 0; i < m-1; i++ {
        a = a.Next
    }
    b := a.Next
    c := b.Next

    for i := 0; i < n-m; i++ {
        t := c.Next
        c.Next = b
        b = c
        c = t
    }

    a.Next.Next = c
    a.Next = b
    return dummyHead.Next
}
```



