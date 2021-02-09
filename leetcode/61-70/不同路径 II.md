63. 不同路径 II

源地址：[63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)

问题描述：

>一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
>
>机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
>
>现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
>
>输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
>输出：2
>解释：
>3x3 网格的正中间有一个障碍物。
>从左上角到右下角一共有 2 条不同的路径：
>1. 向右 -> 向右 -> 向下 -> 向下
>2. 向下 -> 向下 -> 向右 -> 向右
>
>输入：obstacleGrid = [[0,1],[0,0]]
>输出：1
>
>
>提示：
>
>m == obstacleGrid.length
>n == obstacleGrid[i].length
>1 <= m, n <= 100
>obstacleGrid[i][j] 为 0 或 1

``` go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := make([][]int, len(obstacleGrid))
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(obstacleGrid[0]))
    }

    for i := 0; i < len(obstacleGrid); i++ {
        for j := 0; j < len(obstacleGrid[0]); j++ {
            if obstacleGrid[i][j] == 0 {
                if i == 0 && j == 0 {dp[i][j] = 1}
                if i > 0 {dp[i][j] += dp[i-1][j]}
                if j > 0 {dp[i][j] += dp[i][j-1]}
            }
        }
    }

    return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
```



