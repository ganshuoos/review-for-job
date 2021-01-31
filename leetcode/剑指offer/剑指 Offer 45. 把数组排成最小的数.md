剑指 Offer 45. 把数组排成最小的数

地址： [剑指 Offer 45. 把数组排成最小的数](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)

> 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
>
> 
>
>示例 1:
>
>输入: [10,2]
>输出: "102"
>示例 2:
>
>输入: [3,30,34,5,9]
>输出: "3033459"
>
>
>提示:
>
>0 < nums.length <= 100
>说明:
>
>输出结果可能非常大，所以你需要返回一个字符串而不是整数
>拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

``` scala

```

```go
import "strconv"
func minNumber(nums []int) string {
    str := make([]string, len(nums))
    for i := 0; i < len(str); i++ {
        str[i] = strconv.Itoa(nums[i]) 
    }

    quickSort(str, 0, len(nums)-1)
    
    ans := ""
    for _, v := range str {
        ans += string(v)
    }
    return ans
}

func quickSort(str []string, left, right int) {
    if left >= right {return}
    x, i, j := str[left], left-1, right+1
    for i < j {
        i += 1
        for str[i] + x < x + str[i] {
            i += 1
        }

        j -= 1
        for str[j] + x > x + str[j] {
            j -= 1
        }

        if i < j {
            str[i], str[j] = str[j], str[i]
        } 
    }
    quickSort(str, left, j)
    quickSort(str, j+1, right)
}
```

