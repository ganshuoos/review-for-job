剑指 Offer 03. 数组中重复的数字

地址：[剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)

问题描述：

>找出数组中重复的数字。
>
>
>在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
>
>示例 1：
>
>输入：
>[2, 3, 1, 0, 2, 5, 3]
>输出：2 或 3 
>
>
>限制：
>
>2 <= n <= 100000

``` scala
//使用map进行计数 计数不为1 则为重复
import scala.collection.mutable.Map
object Solution {
    def findRepeatNumber(nums: Array[Int]): Int = {
        val map = Map[Int, Int]()

        for (num <- nums) {
            if (map.contains(num) == true) return num
            else {
                map(num) = 1
            }
        }

        return -1
    }
}

/*
map常用方法总结：
1.contains => 是否包含key
2.isEmpty => 判断是否为空
3.head => 头部元素
4.last => 尾部元素
5.max => min =>
6.size 
*/
```

```go
import  "fmt"
func findRepeatNumber(nums []int) int {
    hmap := make(map[int]int)
    ans := make([]int, 0)

    for _, num := range nums {
        if _ , ok := hmap[num]; ok {
            ans = append(ans, num)
        } else {
            hmap[num] = 1
        }
    }

    //fmt.Printf("ans: %v", ans)
    return ans[0]
}
```

