剑指 Offer 59 - I. 滑动窗口的最大值

地址：[剑指 Offer 59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)

> 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
>
> 示例:
>
> 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
> 输出: [3,3,5,5,6,7] 
> 解释: 
>
>   滑动窗口的位置                最大值
> ---------------               -----
> [1  3  -1] -3  5  3  6  7       3
>  1 [3  -1  -3] 5  3  6  7       3
>  1  3 [-1  -3  5] 3  6  7       5
>  1  3  -1 [-3  5  3] 6  7       5
>  1  3  -1  -3 [5  3  6] 7       6
>  1  3  -1  -3  5 [3  6  7]      7
>
>
> 提示：
>
> 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
>

``` 

```

```go
func maxSlidingWindow(nums []int, k int) []int {
    ans := make([]int, 0)
    //单调队列
    queue := make([]int, len(nums)) 

    qHead, qTail := 0, -1
    for i := 0; i < len(nums); i++ {
        if qHead <= qTail && i - k + 1 > queue[qHead] {
            qHead += 1
        }

        for qHead <= qTail && nums[i] >= nums[queue[qTail]] {
            qTail -= 1
        }
        qTail += 1
        queue[qTail] = i

        if i >= k-1 {
            ans = append(ans, nums[queue[qHead]])
        }
    }
    return ans
}
```

