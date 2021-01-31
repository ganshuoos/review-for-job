剑指 Offer 41. 数据流中的中位数

地址：[剑指 Offer 41. 数据流中的中位数](https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/)

>如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
>
>例如，
>
>[2,3,4] 的中位数是 3
>
>[2,3] 的中位数是 (2 + 3) / 2 = 2.5
>
>设计一个支持以下两种操作的数据结构：
>
>void addNum(int num) - 从数据流中添加一个整数到数据结构中。
>double findMedian() - 返回目前所有元素的中位数。
>示例 1：
>
>输入：
>["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
>[[],[1],[2],[],[3],[]]
>输出：[null,null,null,1.50000,null,2.00000]
>示例 2：
>
>输入：
>["MedianFinder","addNum","findMedian","addNum","findMedian"]
>[[],[2],[],[3],[]]
>输出：[null,null,2.00000,null,2.50000]
>
>
>限制：
>
>最多会对 addNum、findMedian 进行 50000 次调用。
>

``` scala

```

```go
//使用对顶堆处理
//leftmax 表示较小侧使用大顶堆
//rightmin 表示较大侧使用小顶堆
//通过两堆堆顶 计算中间值
type MinHeap []int
func (h MinHeap)Len() int {return len(h)}
func (h MinHeap)Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap)Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap)Push (x interface{}) {
    *h = append(*h, x.(int)) 
}
func (h *MinHeap)Pop() interface{} {
    old := *h
    n := len(*h)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MaxHeap []int
func (h MaxHeap)Len() int {return len(h)}
func (h MaxHeap)Less(i, j int) bool {return h[i] > h[j]}
func (h MaxHeap)Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap)Push (x interface{}) {
    *h = append(*h, x.(int)) 
}
func (h *MaxHeap)Pop() interface{} {
    old := *h
    n := len(*h)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MedianFinder struct {
    RightMin *MinHeap
    LeftMax  *MaxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    m := new(MedianFinder)
    m.RightMin = new(MinHeap)
    m.LeftMax  = new(MaxHeap)
    heap.Init(m.RightMin)
    heap.Init(m.LeftMax)
    return *m 
}


func (this *MedianFinder) AddNum(num int)  {
   if this.LeftMax.Len() == this.RightMin.Len() {
       heap.Push(this.RightMin, num)
       heap.Push(this.LeftMax, heap.Pop(this.RightMin))
   } else {
       heap.Push(this.LeftMax, num)
       heap.Push(this.RightMin, heap.Pop(this.LeftMax))
   }
}


func (this *MedianFinder) FindMedian() float64 {

    fmt.Printf("%s: ", "LeftMax")
    for i := 0; i < this.LeftMax.Len(); i++ {
        fmt.Printf("%d, ", (*this.LeftMax)[i])
    }
    fmt.Println()
    fmt.Printf("%s: ", "RightMin")
    for i := 0; i < this.RightMin.Len(); i++ {
        fmt.Printf("%d, ", (*this.RightMin)[i])
    } 
    fmt.Println()

    if (this.LeftMax.Len() + this.RightMin.Len())%2 == 1 {
        return float64((*this.LeftMax)[0])
    } else {
        return float64(((*this.LeftMax)[0] + (*this.RightMin)[0]))/2.0
    }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
```

