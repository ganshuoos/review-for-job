剑指 Offer 25. 合并两个排序的链表

地址：[剑指 Offer 25. 合并两个排序的链表](https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/)

问题描述：

>输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
>
>示例1：
>
>输入：1->2->4, 1->3->4
>输出：1->1->2->3->4->4
>限制：
>
>0 <= 链表长度 <= 1000
>

``` scala
/**
 * Definition for singly-linked list.
 * class ListNode(var _x: Int = 0) {
 *   var next: ListNode = null
 *   var x: Int = _x
 * }
 */
object Solution {
    def mergeTwoLists(la: ListNode, lb: ListNode): ListNode = {
        val dummyHead = new ListNode(0)
        var cur = dummyHead
        var (l1, l2) = (la, lb)
        while (l1 != null && l2 != null) {
            if (l1.x <= l2.x) {
                cur.next = l1
                cur = cur.next
                l1 = l1.next
            } else {
                cur.next = l2
                cur = cur.next
                l2 = l2.next
            }
        }

        while (l1 != null) {
            cur.next = l1
            l1 = l1.next
            cur = cur.next
        }

        while (l2 != null) {
            cur.next = l2
            l2 = l2.next
            cur = cur.next
        }

        return dummyHead.next
    }
}
```

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0, Next: nil}
    cur := dummyHead
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            cur.Next = l1
            cur = cur.Next
            l1 = l1.Next
        } else {
            cur.Next = l2
            cur = cur.Next
            l2 = l2.Next
        }
    }

    for l1 != nil {
        cur.Next = l1
        cur = cur.Next
        l1 = l1.Next
    }

    for l2 != nil {
        cur.Next = l2
        cur = cur.Next
        l2 = l2.Next
    }

    return dummyHead.Next
}
```

