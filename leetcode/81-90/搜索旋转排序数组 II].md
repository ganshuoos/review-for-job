81. 搜索旋转排序数组 II]

源地址：[81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

问题描述：

>假设按照升序排序的数组在预先未知的某个点上进行了旋转。
>
>( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
>
>编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
>
>示例 1:
>
>输入: nums = [2,5,6,0,0,1,2], target = 0
>输出: true
>示例 2:
>
>输入: nums = [2,5,6,0,0,1,2], target = 3
>输出: false
>进阶:
>
>这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
>这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

``` go
func search(nums []int, target int) bool {
    if  len(nums) == 0 {return false}
    left, right, R := 0, len(nums)-1, len(nums)-1
    for R >= 0 && nums[0] == nums[R] {R -= 1}
    if R < 0 {return nums[0] == target}

    left, right = 0, R
    for left < right {
        mid := (left + right + 1) >> 1
        if nums[mid]  >= nums[0] {
            left = mid
        } else {
            right = mid-1
        }
    }

    if target >= nums[0] {
        left, right = 0, left
    } else {
        left, right = left + 1, R 
    }

    for left < right {
        mid := (left + right) >> 1
        if nums[mid] == target {
            return true
        } else if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return nums[right] == target
}
```



