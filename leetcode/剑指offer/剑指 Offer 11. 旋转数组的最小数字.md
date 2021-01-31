剑指 Offer 11. 旋转数组的最小数字

地址：[剑指 Offer 11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

问题描述：

>把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  
>
>示例 1：
>
>输入：[3,4,5,1,2]
>输出：1
>示例 2：
>
>输入：[2,2,2,0,1]
>输出：0
>注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

``` scala
object Solution {
    def minArray(nums: Array[Int]): Int = {
        var left = 0
        var right = nums.length - 1

        while (left < right) {
            val mid = (left + right)/2
            if (nums(mid) < nums(right)) {
                right = mid
            } else if (nums(mid) > nums(right)) {
                left = mid + 1
            } else {
                right -= 1
            }
        }

        return nums(right)
    }
}
```

```go
func minArray(nums []int) int {
    left := 0
    right := len(nums) - 1

    for (left < right) {
        mid := (left + right)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else if nums[mid] < nums[right] {
            right = mid 
        } else {
            right -= 1
        }
    }

    return nums[right]
}
```

