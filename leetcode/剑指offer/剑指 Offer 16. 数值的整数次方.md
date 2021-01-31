剑指 Offer 16. 数值的整数次方

地址：[剑指 Offer 16. 数值的整数次方](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

问题描述：

>实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
>
> 
>
>示例 1:
>
>输入: 2.00000, 10
>输出: 1024.00000
>示例 2:
>
>输入: 2.10000, 3
>输出: 9.26100
>示例 3:
>
>输入: 2.00000, -2
>输出: 0.25000
>解释: 2-2 = 1/22 = 1/4 = 0.25
>
>
>说明:
>
>-100.0 < x < 100.0
>n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

``` scala
object Solution {
    def myPow(x: Double, n: Int): Double = {
        var eNum: Long = n
        var base = x

        if (n == 0) return 1
        if (n == 1) return x
        if (n < 0) {
            base = 1/base
            eNum = -eNum
        }
		
        //使用递归思路 将奇数分为 (base)n/2 (base)n/2 base
        //偶数分为(base)n/2 (base)n/2
        var res = myPow(base, (eNum >> 1).toInt)
        res *= res

        if ((eNum & 1) == 1) res *= base

        return res
    }
}
```

```go
func myPow(x float64, n int) float64 {
    base := x
    eNum := n 

    if eNum == 0 {return 1}
    if eNum == 1 {return base}

    if eNum < 0 {
        base = 1/base
        eNum = -eNum
    }

    res := myPow(base, eNum >> 1)
    res *= res

    if (eNum & 1) == 1 {res *= base}

    return res
}
```

