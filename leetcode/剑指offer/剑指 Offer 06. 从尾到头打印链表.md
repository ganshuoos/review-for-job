剑指 Offer 06. 从尾到头打印链表

地址：[剑指 Offer 06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

问题描述：

>输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
>
> 
>
>示例 1：
>
>输入：head = [1,3,2]
>输出：[2,3,1]
>
>
>限制：
>
>0 <= 链表长度 <= 10000
>

``` scala
/**
 * Definition for singly-linked list.
 * class ListNode(var _x: Int = 0) {
 *   var next: ListNode = null
 *   var x: Int = _x
 * }
 */
//存储结果 然后翻转
//也可以使用栈
//也可以翻转链表
// pre cur = head
// temp = cur.next
// cur.next = pre
// pre = cur
// cur = cur.next
object Solution {
    def reversePrint(head: ListNode): Array[Int] = {
        var newHead = head
        var ans = new Array[Int](0)
        if (head == null) return ans

        while (newHead != null) {
            ans = ans.prepended(newHead.x)
            newHead = newHead.next
            //ans = ans.prepended(newHead.x)
        }
        return ans
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
//go stack 利用slice实现
func reversePrint(head *ListNode) []int {
    if head == nil {return nil}

    res := make([]int, 0)

    for (head != nil) {
        res = append(res, head.Val)
        head = head.Next
    }

    for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
        temp := res[i]
        res[i] = res[j]
        res[j] = temp
    } 

    return res
}
```

