剑指 Offer 22. 链表中倒数第k个节点

地址：[剑指 Offer 22. 链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)

问题描述：

>输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。
>
> 
>
>示例：
>
>给定一个链表: 1->2->3->4->5, 和 k = 2.
>
>返回链表 4->5.
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
    def getKthFromEnd(head: ListNode, k: Int): ListNode = {
        var (fast, slow) = (head, head)
        var num = k
        while (num > 0) {
            fast = fast.next
            num -= 1
        }
        while (fast != null) {
            fast = fast.next
            slow = slow.next
        }
        return slow 
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
import "fmt"
func getKthFromEnd(head *ListNode, k int) *ListNode {
    //dummyHead := &ListNode{Val:0, Next:head}
    fast := head
    for k > 0 {
        //if fast.Next != nil {fast = fast.Next}
        fast = fast.Next
        k -= 1
    }

    for fast != nil {
        fast = fast.Next
        head = head.Next
    }

    return head
}
```

