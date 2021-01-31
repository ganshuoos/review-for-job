剑指 Offer 10- I. 斐波那契数列

地址：[剑指 Offer 10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

问题描述：

>写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
>
>F(0) = 0,   F(1) = 1
>F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
>斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
>
>答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 
>
>示例 1：
>
>输入：n = 2
>输出：1
>示例 2：
>
>输入：n = 5
>输出：5
>
>
>提示：
>
>0 <= n <= 100
>

``` scala
object Solution {
    def fib(n: Int): Int = {
        if (n == 0) return 0
        if (n == 1) return 1

        var fibArr = Array.fill(n+1)(0)
        fibArr(0) = 0
        fibArr(1) = 1
        for (i <- 2 to n) {
            fibArr(i) = fibArr(i-1) + fibArr(i-2)
            fibArr(i) = fibArr(i) % 1000000007
        }
        return fibArr(n)
    }
}
```

```go
func fib(n int) int {
    f0, f1 := 0, 1
	for i := 0; i < n; i++ {
        f0,f1 = f1,(f0 + f1) % 1000000007;
	}
	return f0
}
```

