剑指 Offer 42. 连续子数组的最大和

地址：[剑指 Offer 42. 连续子数组的最大和](https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/)

>输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
>
>要求时间复杂度为O(n)。
>
> 
>
>示例1:
>
>输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
>输出: 6
>解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
>
>
>提示：
>
>1 <= arr.length <= 10^5
>-100 <= arr[i] <= 100

``` scala

```

```go
func maxSubArray(nums []int) int {
    length := len(nums)
    if length == 0 {return 0}
    if length == 1 {return nums[0]}
    res := nums[0]
    dp := make([]int, length)
    dp[0] = nums[0]
    for i := 1; i < length; i++ {
        dp[i] = max(dp[i-1]+nums[i], nums[i])
        res = max(res, dp[i])
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```

