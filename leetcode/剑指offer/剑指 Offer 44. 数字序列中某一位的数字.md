剑指 Offer 44. 数字序列中某一位的数字

地址：[剑指 Offer 44. 数字序列中某一位的数字](https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/)

>数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
>
>请写一个函数，求任意第n位对应的数字。
>
> 
>
>示例 1：
>
>输入：n = 3
>输出：3
>示例 2：
>
>输入：n = 11
>输出：0
>
>
>限制：
>
>0 <= n < 2^31
>

``` scala

```

```go
import "strconv"

func findNthDigit(n int) int {
    //start表开始位置， dight表数位， count表个数
    start, digit, count := 1, 1, 9
    //不断去除各个数位的数的个数，确定数位！
    for n > count * digit {
        n -= count * digit
        digit += 1
        start *= 10
        count *= 10
    }

    //确定是当前数位下的第几个数， 由于每个int数可拆为digit个1位数
    //所以 n / k 取整表示是从当前start开始的第几个数
    start += (n + digit - 1)/digit - 1

    //确定是当前这个int数的第几个数位， 当前每个int数的digit是确定的，只要对%digit进行取余，就知道其位于第几个位置
    if n % digit == 0 {
        n = digit
    } else {
        n = n % digit
    }

    ans := int((strconv.Itoa(start))[n-1] - '0')
    return ans
}
```

