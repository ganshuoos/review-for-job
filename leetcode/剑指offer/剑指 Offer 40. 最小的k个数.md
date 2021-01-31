剑指 Offer 40. 最小的k个数

地址：[剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

>输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
>
> 
>
>示例 1：
>
>输入：arr = [3,2,1], k = 2
>输出：[1,2] 或者 [2,1]
>示例 2：
>
>输入：arr = [0,1,2,1], k = 1
>输出：[0]
>
>
>限制：
>
>0 <= k <= arr.length <= 10000
>0 <= arr[i] <= 10000

``` scala

```

```go
//排序
import "sort"
func getLeastNumbers(arr []int, k int) []int {
    sort.Ints(arr)
    return arr[:k]
}

//堆方法
import "container/heap"

type intHeap []int

func (h intHeap)Len() int {return len(h)}
func (h intHeap)Less (i, j int) bool {return h[i] > h[j]}
func (h intHeap)Swap (i, j int) {h[i], h[j] = h[j], h[i]} 
func (h *intHeap)Push (x interface{}) {(*h) = append(*h, x.(int))}
func (h *intHeap)Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func getLeastNumbers(arr []int, k int) []int {
    h := &intHeap{}
    heap.Init(h)
    for _, v := range arr {
        heap.Push(h, v)
        for h.Len() > k {
            _ = heap.Pop(h)
        }
    }
    return (*h)[:k]
}

//注：
//go中最小堆， 即func(h heap)Less(i, j int) bool {return h[i] < h[j]}
//堆中元素存放 由小至大， Pop操作弹出堆中最小值（堆顶）
//&main.intHeap{1, 2, 3}
//&main.intHeap{2, 3}
//&main.intHeap{2, 3, 4}
//&main.intHeap{3, 4}
//反之
//&main.intHeap{3, 1, 2}
//&main.intHeap{2, 1}
//&main.intHeap{4, 1, 2}
//&main.intHeap{2, 1}
```

