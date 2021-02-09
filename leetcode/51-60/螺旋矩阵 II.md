59. 螺旋矩阵 II

源地址：[59. 螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

问题描述：

>给你一个正整数 `n` ，生成一个包含 `1` 到 `n2` 所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形矩阵 `matrix` 。
>
>输入：n = 3
>输出：[[1,2,3],[8,9,4],[7,6,5]]
>示例 2：
>
>输入：n = 1
>输出：[[1]]
>
>
>提示：
>
>1 <= n <= 20

``` go
func generateMatrix(n int) [][]int {
     dx := []int{0, 1, 0, -1}
     dy := []int{1, 0, -1, 0}
     res := make([][]int, n)
     for i := 0; i < n; i++ {
         res[i] = make([]int, n)
     }

     st := make([][]bool, n)
     for i := 0; i < n; i++ {
         st[i] = make([]bool, n)
     }

     for i, x, y, d := 1, 0, 0, 0; i <= n*n; i++ {
         res[x][y] = i 
         st[x][y] = true
         a, b := x + dx[d], y + dy[d]

         if a < 0 || a >= n || b < 0 || b >= n || st[a][b] == true {
             d = (d+1)%4
             a, b = x + dx[d], y + dy[d]
         }

         x, y = a, b  
     }

     return res
}
```



