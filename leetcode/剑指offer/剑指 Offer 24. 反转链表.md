剑指 Offer 24. 反转链表

地址：[剑指 Offer 24. 反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)

问题描述：

>定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
>
> 
>
>示例:
>
>输入: 1->2->3->4->5->NULL
>输出: 5->4->3->2->1->NULL
>
>
>限制：
>
>0 <= 节点个数 <= 5000
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
    def reverseList(head: ListNode): ListNode = {
        if (head == null) return head

        val dummyHead = ListNode(0)
        dummyHead.next = head
        var newHead = head
        var tail = head
        while (newHead.next != null) {
            val temp = newHead.next
            newHead.next = temp.next
            dummyHead.next = temp
            temp.next = tail
            tail = temp
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
func reverseList(head *ListNode) *ListNode {
    if head == nil {return head}
    dummyHead := &ListNode{Val: 0 , Next: head}
    tail := head
    for head.Next != nil {
        temp := head.Next
        head.Next = temp.Next
        dummyHead.Next = temp
        temp.Next = tail
        tail = temp  
    }
    return dummyHead.Next
}
```

