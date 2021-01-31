剑指 Offer 13. 机器人的运动范围

地址：[剑指 Offer 13. 机器人的运动范围](https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/)

问题描述：

>地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
>
> 
>
>示例 1：
>
>输入：m = 2, n = 3, k = 1
>输出：3
>示例 2：
>
>输入：m = 3, n = 1, k = 0
>输出：1
>提示：
>
>1 <= n,m <= 100
>0 <= k <= 20

``` scala
//DFS
object Solution {
    def movingCount(m: Int, n: Int, k: Int): Int = {
        val visited = Array.fill(m+1, n+1)(false)
        return dfs(visited, m, n, k, 0, 0) 
    }

    def dfs(visited: Array[Array[Boolean]], m: Int, n: Int, k: Int, i: Int, j: Int): Int = {
        if (i < 0 || i >= m || j < 0 || j >= n || visited(i)(j) == true || sum(i) + sum(j) > k) return 0
        visited(i)(j) = true
        return 1 + dfs(visited, m, n, k, i+1, j) + dfs(visited, m, n, k, i, j+1)
    }

    def sum(x: Int): Int = {
        var sum = 0
        var tempX = x
        while (tempX > 0) {
            sum += tempX % 10
            tempX /= 10
        }
        return sum
    }
}


//BFS
import scala.collection.mutable.Queue
import util.control.Breaks._
object Solution {
    def movingCount(m: Int, n: Int, k: Int): Int = {
        val queue = Queue[(Int, Int)]()
        val visited = Array.fill(m, n)(false)
        var res = 0

        queue.enqueue((0, 0))
        while(queue.size > 0) {
            val a = queue.dequeue
            val x = a._1
            val y = a._2
            
            breakable{
                if (x < 0 || x >= m || y < 0 || y >= n || visited(x)(y) == true || sum(x) + sum(y) > k) break()
                visited(x)(y) = true
                res += 1
                queue.enqueue((x+1, y))
                queue.enqueue((x, y+1))
            }
        }
        return res
    }

    def sum(i: Int): Int = {
        var sum = 0
        var tempI = i
        while (tempI > 0) {
            sum += tempI%10
            tempI /= 10
        }
        return sum
    }
}
```

```go
//DFS
func movingCount(m int, n int, k int) int {
    //visited := [m][n]bool{}
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    return dfs(visited, m, n, k, 0, 0)
}

func dfs(visited [][]bool, m, n, k, i, j int) int {
    if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == true || sum(i) + sum(j) > k { return 0}
    visited[i][j] = true
    return 1 + dfs(visited, m, n, k, i+1, j) + dfs(visited, m, n, k, i, j+1)
} 

func sum(x int) int {
    res := 0
    for (x > 0) {
        res += x % 10
        x /= 10
    }
    return res
}

//BFS
import "container/list"
func movingCount(m int, n int, k int) int {
    res := 0
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    queue := list.New()
    queue.PushBack(0)
    queue.PushBack(0)

    for (queue.Len() > 0) {
        temp := queue.Front()
        x := temp.Value.(int)
        queue.Remove(temp)
        temp = queue.Front()
        y := temp.Value.(int)
        queue.Remove(temp)

        if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] == true ||sum(x) + sum(y) > k {continue}
        visited[x][y] = true
        res += 1
        queue.PushBack(x+1)
        queue.PushBack(y)
        queue.PushBack(x)
        queue.PushBack(y+1) 
    }

    return res
}

func sum(x int) int {
    res := 0 
    for (x > 0) {
        res += x%10
        x /= 10
    }
    return res
}
```

