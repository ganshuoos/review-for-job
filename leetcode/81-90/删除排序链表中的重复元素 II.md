82. 删除排序链表中的重复元素 II

源地址：[82. 删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/)

问题描述：

>给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
>
>示例 1:
>
>输入: 1->2->3->3->4->4->5
>输出: 1->2->5
>示例 2:
>
>输入: 1->1->1->2->3
>输出: 2->3

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    dummyHead := &ListNode{-1, head}
    p := dummyHead
    for p.Next != nil {
        q := p.Next.Next
        for q != nil && q.Val == p.Next.Val {q = q.Next}
        if p.Next.Next == q {
            p = p.Next
        } else {
            p.Next = q
        }
    }
    return dummyHead.Next
}
```



