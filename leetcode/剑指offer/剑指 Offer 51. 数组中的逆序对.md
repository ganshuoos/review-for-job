剑指 Offer 51. 数组中的逆序对

地址：[剑指 Offer 51. 数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)

问题描述：

>在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
>
> 
>
>示例 1:
>
>输入: [7,5,6,4]
>输出: 5
>
>
>限制：
>
>0 <= 数组长度 <= 50000
>

``` scala
object Solution {
    val tmp = new Array[Int](50005)

    def reversePairs(nums: Array[Int]): Int = {
        return mergeSort(nums, 0, nums.length-1).toInt
    }

    def mergeSort(nums: Array[Int], left: Int, right: Int): Long = {
        if (left >= right) {return 0}
        val mid = (left + right) >> 1
        var res: Long = mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)

        var k = 0
        var i = left
        var j = mid + 1
        while (i <= mid && j <= right) {
            if (nums(i) <= nums(j)) {
                tmp(k) = nums(i)
                k += 1
                i += 1
            } else {
                tmp(k) = nums(j)
                res = res + mid - i + 1
                k += 1
                j += 1 
            }
        }

        while (i <= mid) {
            tmp(k) = nums(i)
            k += 1
            i += 1
        }

        while (j <= right) {
            tmp(k) = nums(j)
            k += 1
            j += 1 
        }

        i = left
        j = 0

        while (i <= right) {
            nums(i) = tmp(j)
            i += 1
            j += 1
        }

        return res
    }
}
```

```go
var tmp []int
func reversePairs(nums []int) int {
    tmp = make([]int, 50005)
    return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
    if left >= right {return 0}

    mid := (left + right) >> 1
    res := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)

    i, j, k := left, mid+1, 0
    for  i <= mid && j <= right {
        if nums[i] <= nums[j] {
            tmp[k] = nums[i]
            k += 1
            i += 1
        } else {
            res += mid - i + 1
            tmp[k] = nums[j]
            k += 1
            j += 1
        }
    }

    for i <= mid {
        tmp[k] = nums[i]
        k += 1
        i += 1
    }

    for j <= right {
        tmp[k] = nums[j]
        k += 1
        j += 1
    }

    for i, j := left, 0; i <= right; {
        nums[i] = tmp[j]
        i += 1
        j += 1
    }

    return res
}
```

