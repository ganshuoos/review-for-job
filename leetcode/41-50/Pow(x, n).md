50. Pow(x, n)

源地址：[50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)

问题描述：

>实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
>
> 
>
>示例 1：
>
>输入：x = 2.00000, n = 10
>输出：1024.00000
>示例 2：
>
>输入：x = 2.10000, n = 3
>输出：9.26100
>示例 3：
>
>输入：x = 2.00000, n = -2
>输出：0.25000
>解释：2-2 = 1/22 = 1/4 = 0.25
>
>
>提示：
>
>-100.0 < x < 100.0
>-231 <= n <= 231-1
>-104 <= xn <= 104

``` go
func myPow(x float64, n int) float64 {
    minus := false
    if n < 0 {minus = true}
    res := 1.0

    for i := abs(n); i > 0; i >>= 1 {
        if (i & 1) > 0 {
            res *= x
        }
        x *= x
    }

    if minus == true {return 1/res}
    return res 
}

func abs(x int) int {
    if x < 0 {return -x}
    return x
}
```



