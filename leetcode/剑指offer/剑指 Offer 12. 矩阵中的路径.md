剑指 Offer 12. 矩阵中的路径

地址：[剑指 Offer 12. 矩阵中的路径](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)

问题描述：

>请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
>
>[["a","b","c","e"],
>["s","f","c","s"],
>["a","d","e","e"]]
>
>但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
>
> 
>
>示例 1：
>
>输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
>输出：true
>示例 2：
>
>输入：board = [["a","b"],["c","d"]], word = "abcd"
>输出：false
>提示：
>
>1 <= board.length <= 200
>1 <= board[i].length <= 200
>注意：本题与主站 79 题相同：https://leetcode-cn.com/problems/word-search/

``` scala
import util.control.Breaks._

object Solution {
    val dx = Array[Int](-1, 0, 1, 0)
    val dy = Array[Int](0, -1, 0, 1)

    def exist(board: Array[Array[Char]], word: String): Boolean = {
        for (i <- 0 to board.length-1) {
            for (j <- 0 to board(0).length - 1) {
                if (dfs(board, word, 0, i, j) == true) return true
            }
        }

        return false
    }

    def dfs(board: Array[Array[Char]], word: String, u: Int, i: Int, j: Int): Boolean = {
        if (board(i)(j) != word(u)) {return false}
        if (u == word.length-1) return true

        val temp = board(i)(j)
        board(i)(j) = '.'

        for (k <- 0 to 3) {
            breakable{
                var a = i + dx(k)
                var b = j + dy(k)

                if (a < 0 || a >= board.length || b < 0 || b >= board(0).length || board(a)(b) == '.') break()
                if (dfs(board, word, u+1, a, b) == true) return true
            }
        }
        board(i)(j) = temp
        return false
    }
}
```

```go
var dx = []int{-1, 0, 1, 0}
var dy = []int{0, -1, 0, 1}

func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(board, word, 0, i, j) == true {return true}
        }
    } 
    return false
}

func dfs(board [][]byte, word string, u, i, j int) bool {
    if board[i][j] != word[u] {return false}
    if u == len(word)-1 {return true}
 
    temp := board[i][j]
    board[i][j] = '.'
    for k := 0; k < 4; k++ {
        a, b := i+dx[k], j+dy[k]
        if (a < 0 || a >= len(board) || b < 0|| b >= len(board[0])|| board[a][b] == '.' ) {continue}
        if (dfs(board, word, u+1, a, b) == true) {return true}
    }
    board[i][j] = temp
    return false
}
```

