23. 合并K个升序链表

源地址：[23. 合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

问题描述：

>给你一个链表数组，每个链表都已经按升序排列。
>
>请你将所有链表合并到一个升序链表中，返回合并后的链表。
>
> 
>
>示例 1：
>
>输入：lists = [[1,4,5],[1,3,4],[2,6]]
>输出：[1,1,2,3,4,4,5,6]
>解释：链表数组如下：
>[
>  1->4->5,
>  1->3->4,
>  2->6
>]
>将它们合并到一个有序链表中得到。
>1->1->2->3->4->4->5->6
>示例 2：
>
>输入：lists = []
>输出：[]
>示例 3：
>
>输入：lists = [[]]
>输出：[]
>
>
>提示：
>
>k == lists.length
>0 <= k <= 10^4
>0 <= lists[i].length <= 500
>-10^4 <= lists[i][j] <= 10^4
>lists[i] 按 升序 排列
>lists[i].length 的总和不超过 10^4

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "container/heap"
type minHeap []*ListNode
func (h minHeap) Len() int { return len(h)}
func (h minHeap) Less(i, j int) bool {return h[i].Val < h[j].Val}
func (h minHeap) Swap(i, j int)  {h[i], h[j] = h[j], h[i]}
func (h *minHeap) Push(x interface{}) {*h = append(*h, x.(*ListNode))}
func (h *minHeap) Pop() interface{} {
    temp := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return temp 
}


func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}

    dummyHead := &ListNode{}
    head := dummyHead

    h := minHeap{}
    heap.Init(&h)

    for _, v := range lists {
        if v != nil {
            heap.Push(&h, v)
        }
    }

    for h.Len() > 0 {
        min :=  heap.Pop(&h).(*ListNode)
        head.Next = min
        head = head.Next
        if min.Next != nil {heap.Push(&h, min.Next)}
    }
    return dummyHead.Next
}
```



