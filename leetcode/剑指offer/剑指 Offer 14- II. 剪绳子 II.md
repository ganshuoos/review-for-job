剑指 Offer 14- II. 剪绳子 II

地址：[剑指 Offer 14- II. 剪绳子 II](https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/)

问题描述：

>给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
>
>答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 
>
>示例 1：
>
>输入: 2
>输出: 1
>解释: 2 = 1 + 1, 1 × 1 = 1
>示例 2:
>
>输入: 10
>输出: 36
>解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
>
>
>提示：
>
>2 <= n <= 1000
>注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/
>
>来源：力扣（LeetCode）
>链接：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof
>著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

``` scala
object Solution {
    def cuttingRope(n: Int): Int = {
       var num = n
       var p: Long = 1
       if (n <= 3) return 1*(n-1)
       while (num >= 5) {
           num -= 3
           p = p * 3%1000000007
       }
       
       return ((p * num)%1000000007).toInt
    }
}
```

```go
import "fmt"
func cuttingRope(n int) int {
    if n <= 3 {return 1 * (n-1)}
    var p int64 = 1
    for (n >= 5) {
        n -= 3
        p = (p * 3)%1000000007 
        
    }
    return (int)(p * int64(n) % 1000000007)
}
```

