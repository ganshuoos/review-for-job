剑指 Offer 53 - I. 在排序数组中查找数字 I

地址：[剑指 Offer 53 - I. 在排序数组中查找数字 I](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)

> 统计一个数字在排序数组中出现的次数。
>
>  
>
> 示例 1:
>
> 输入: nums = [5,7,7,8,8,10], target = 8
> 输出: 2
> 示例 2:
>
> 输入: nums = [5,7,7,8,8,10], target = 6
> 输出: 0
>
>
> 限制：
>
> 0 <= 数组长度 <= 50000
>

``` 

```

```go
func search(nums []int, target int) int {
    //排除空值
    if len(nums) == 0 {return 0}
    left, right := 0, len(nums)-1
    first := searchLeft(nums, left, right, target)
    last  := searchRight(nums, left, right, target)
    //用于排除target值不在区间内
    if nums[first] != target {return 0}
    return last - first + 1
}

func searchLeft(nums []int, left, right, target int) int {
    for left < right {
        mid := (left + right) >> 1
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func searchRight(nums []int, left, right, target int) int {
    for left < right {
        mid := (left + right +1) >> 1
        if nums[mid] <= target {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return left
}
```

