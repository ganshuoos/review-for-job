51. N 皇后

源地址：[51. N 皇后](https://leetcode-cn.com/problems/n-queens/)

问题描述：

>n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
>
>给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
>
>每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
>
>输入：n = 4
>输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
>解释：如上图所示，4 皇后问题存在两个不同的解法。
>示例 2：
>
>输入：n = 1
>输出：[["Q"]]
>
>
>提示：
>
>1 <= n <= 9
>皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

``` go
var col []bool
var dg  []bool
var udg []bool
var queue []int
var res [][]string

func solveNQueens(n int) [][]string {
    col = make([]bool, n)
    dg = make([]bool, n*2)
    udg = make([]bool, n*2)
    queue = make([]int, n)
    res = make([][]string, 0)
    dfs(0, n)
    return res 
}

func dfs(u, n int) {
    if u == n {
        borad := generateBoard(n)
        res = append(res, borad)
        return
    }

    for i := 0; i < n; i++ {
        if col[i] == false && udg[u+i] == false && dg[u-i+n] == false {
            col[i], udg[u+i], dg[u-i+n] = true, true, true
            queue[u] = i
            dfs(u+1, n)
            queue[u] = 0
            col[i], udg[u+i], dg[u-i+n] = false, false, false 
        } 
    }
}

func generateBoard(n int) []string {
    board := make([]string, 0)
    for i := 0; i < n; i++ {
        row := make([]byte, n)
        for j := 0; j < n; j++ {
            row[j] = '.'
        }
        row[queue[i]] = 'Q'
        board = append(board, string(row))
    }
    return board
}
```



