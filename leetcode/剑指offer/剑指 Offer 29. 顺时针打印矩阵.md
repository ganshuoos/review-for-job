剑指 Offer 29. 顺时针打印矩阵

地址：[剑指 Offer 29. 顺时针打印矩阵](https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)

>输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
>
> 
>
>示例 1：
>
>输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
>输出：[1,2,3,6,9,8,7,4,5]
>示例 2：
>
>输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
>输出：[1,2,3,4,8,12,11,10,9,5,6,7]
>
>
>限制：
>
>0 <= matrix.length <= 100
>0 <= matrix[i].length <= 100
>

``` scala

```

```go
func spiralOrder(matrix [][]int) []int {
    res := make([]int, 0)
    if len(matrix) == 0 || len(matrix[0]) == 0 {return res} 
    row, col := len(matrix), len(matrix[0])


    var dx = []int{0, 1, 0, -1}
    var dy = []int{1, 0, -1, 0}

    st := make([][]bool, row)
    for i := 0; i < len(st); i++ {
        st[i] = make([]bool, col)
    }

    for i, x, y, d := 0, 0, 0, 0; i < row * col; i++ {
        res = append(res, matrix[x][y])
        st[x][y] = true

        a, b := x + dx[d], y + dy[d]
        if a < 0 || a >= row || b < 0 || b >= col || st[a][b] == true {
            d = (d + 1) % 4
            a, b = x + dx[d], y + dy[d]
        }
        x, y = a, b
    }

    return res
}
```

