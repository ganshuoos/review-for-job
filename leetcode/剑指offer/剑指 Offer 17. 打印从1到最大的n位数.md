剑指 Offer 17. 打印从1到最大的n位数

地址：[剑指 Offer 17. 打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

问题描述：

>输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
>
>示例 1:
>
>输入: n = 1
>输出: [1,2,3,4,5,6,7,8,9]
>
>
>说明：
>
>用返回一个整数列表来代替打印
>n 为正整数

``` scala
import scala.collection.mutable.ListBuffer
object Solution {
    def printNumbers(n: Int): Array[Int] = {
        val number = Array.fill(n)("0")
        val res = ListBuffer[Int]()
		
        //利用数组 模拟数 构成全排列
        for (i <- 0 to 9) {
            number(0) = i.toString
            dfs(res, number, n, 0)
        }
        
        //把第一个零去掉
        return res.drop(1).toArray
    }

    def dfs(res: ListBuffer[Int], number: Array[String], length: Int, index: Int) {
        //铺满数组
        if (index == length-1) {
            val ans = number.mkString.toInt
            res.append(ans)
            return
        }
        
        for(i <- 0 to 9) {
            number(index+1) = i.toString
            dfs(res, number, length, index+1)
        }

    }
}
```

```go
import "fmt"
var charNum = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var res []int

func printNumbers(n int) []int {
    num := make([]byte, n)
    res = make([]int, 0)

    for _, v := range charNum {
        num[0] = v
        dfs(num, n, 0)
    }

    return res[1:]
}

func dfs(num []byte, length, index int) {
    if index == length - 1 {
        val, _ := strconv.Atoi(string(num))
        //fmt.Printf("val: %d\n", val)
        //if err != nil {
        //    res = append(res, val)
        //}
        res = append(res, val)
        //fmt.Printf("res: %v\n", res)
        return 
    }

    for _, v := range charNum {
        num[index + 1] = v
        dfs(num, length, index+1)
    }
}
```

