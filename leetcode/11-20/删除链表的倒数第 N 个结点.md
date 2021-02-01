19. 删除链表的倒数第 N 个结点

源地址：[19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

问题描述：

>给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。
>
>**进阶：**你能尝试使用一趟扫描实现吗？
>
>输入：head = [1,2,3,4,5], n = 2
>输出：[1,2,3,5]
>示例 2：
>
>输入：head = [1], n = 1
>输出：[]
>示例 3：
>
>输入：head = [1,2], n = 1
>输出：[1]
>
>
>提示：
>
>链表中结点的数目为 sz
>1 <= sz <= 30
>0 <= Node.val <= 100
>1 <= n <= sz

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{0, head}
    pre, cur := dummyHead, dummyHead
    for i := 0; i < n; i++ {
        pre = pre.Next
    }
    for pre.Next != nil{
        pre = pre.Next
        cur = cur.Next
    }

    cur.Next = cur.Next.Next
    return dummyHead.Next
}
```



