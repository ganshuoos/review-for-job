剑指 Offer 47. 礼物的最大价值

地址： [剑指 Offer 47. 礼物的最大价值](https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/)

> 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
>
>  
>
> 示例 1:
>
> 输入: 
> [
>   [1,3,1],
>   [1,5,1],
>   [4,2,1]
> ]
> 输出: 12
> 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
>
>
> 提示：
>
> 0 < grid.length <= 200
> 0 < grid[0].length <= 200

``` scala

```

```go
func maxValue(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])

    dp := make([][]int, row+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, col+1)
    }

    for i := 1; i <= row; i++ {
        for j := 1; j <= col; j++ {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
        }
    }

    return dp[row][col]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```

