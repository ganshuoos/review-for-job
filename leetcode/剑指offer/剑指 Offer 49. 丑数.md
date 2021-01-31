剑指 Offer 49. 丑数

地址：[剑指 Offer 49. 丑数](https://leetcode-cn.com/problems/chou-shu-lcof/)

> 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
>
>  
>
> 示例:
>
> 输入: n = 10
> 输出: 12
> 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
> 说明:  
>
> 1 是丑数。
> n 不超过1690。

``` scala

```

```go
func nthUglyNumber(n int) int {
    res := make([]int, 0)
    two  := 0
    three:= 0
    five := 0
    res = append(res, 1)

    if n == 1 {return 1}
    for len(res) < n {
        ans := min(2*res[two], min(3*res[three], 5*res[five]))
        if ans == 2*res[two] {two += 1}
        if ans == 3*res[three] {three += 1}
        if ans == 5*res[five] {five += 1}
        res = append(res, ans)
    }
    return res[len(res)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
```

