119. 杨辉三角 II

源地址：[119. 杨辉三角 II](https://leetcode-cn.com/problems/pascals-triangle-ii/)

问题描述：

>给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
>
>
>
>在杨辉三角中，每个数是它左上方和右上方的数的和。
>
>示例:
>
>输入: 3
>输出: [1,3,3,1]
>进阶：
>
>你可以优化你的算法到 O(k) 空间复杂度吗？
>

``` go
func getRow(n int) []int {
    res := make([][]int, 0)
    for i := 0; i <= n; i++ {
        floor := make([]int, i+1)
        floor[0], floor[i] = 1, 1
        for j := 1; j < i; j++ {floor[j] = res[i-1][j-1] + res[i-1][j]}
        res = append(res, floor)
    }
    return res[n]
}
```



