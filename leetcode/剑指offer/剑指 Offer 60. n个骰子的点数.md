剑指 Offer 60. n个骰子的点数

地址：[剑指 Offer 60. n个骰子的点数](https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/)

> 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
>
>  
>
> 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
>
>  
>
> 示例 1:
>
> 输入: 1
> 输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
> 示例 2:
>
> 输入: 2
> 输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
>
>
> 限制：
>
> 1 <= n <= 11

``` 

```

```go
import (
    "fmt"
    "math"
)

func dicesProbability(n int) []float64 {
    total := math.Pow(float64(6),float64(n))
    
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, 6*n+1)
    }

    //初始化
    for s := 0; s <= 6; s++ {
        dp[1][s] = 1
    }

    for i := 2; i <= n; i++ {
        for s := i; s <= 6*i; s++ {
            for d := 1; d <= 6; d++ {
                //剪枝 第i次骰子结果为s 第i次骰子结果为1~6
                //第i-1次骰子最少为i-1, 
                if s-d < i-1 {
                    break
                } else {
                    dp[i][s] += dp[i-1][s-d]
                }
            }
        }
    }

    res := make([]float64, 0)
    for i := n; i <= 6*n; i++ {
        res = append(res, float64(dp[n][i])/total)
    }
    return res
}
```

