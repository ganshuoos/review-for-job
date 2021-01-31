剑指 Offer 43. 1～n 整数中 1 出现的次数

地址：[剑指 Offer 43. 1～n 整数中 1 出现的次数](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/)

>输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
>
>例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
>
> 
>
>示例 1：
>
>输入：n = 12
>输出：5
>示例 2：
>
>输入：n = 13
>输出：6
>
>
>限制：
>
>1 <= n < 2^31

``` scala

```

```go
import "fmt"

func countDigitOne(n int) int {
    nums := make([]int, 0)
    res := 0
    for n > 0 {
        nums = append(nums, n%10)
        n /= 10
    }
    nums = reverse(nums)
    for i := range nums {
        //取当前位置为mid
        mid := nums[i]
        //使用mid分割left， right
        left, right, p := 0, 0, 1
        //计算左侧值
        for j := 0; j < i; j++ {
            left = left * 10 + nums[j]
        }
        //计算右侧值及剩余部分次数
        for j := i+1; j < len(nums); j++ {
            right = right * 10 + nums[j]
            p = p * 10
        }
        
        if mid == 0 {
        //当mid == 0时， 前缀为（0 -> left-1）的数存在（0 -> left -1）1 （right）
        //对于当前位为1的个数统计， 只受前缀影响（right可取任意数）， 共left个， p为10进制位置，故mid为1的个数共有left * p  
        //前缀为left时， mid位为0， 不存在mid为1的个数
            res = res + left * p
        } else if mid == 1 {
        //当mid == 1时， 前缀为（0 -> left）的数存在（0 -> left ）1 （right）
        //对于当前位为0时, 由上可得 left * p  
        //对于当前位为1时， right任意值都可以放入及自身的1次放入统计
            res = res + left * p + 1 + right
        } else {
        //当mid > 2时， 高于1时，都不需考虑
        //与mid == 1时不同点在于， 1的right侧可以取任意值， 故直接在left上加1， 乘以10位数即可
            res = res + (left + 1) * p
        }
    }
    return res
}

//旋转
func reverse(nums []int) []int {
   length := len(nums)
   for i, j := 0, length-1; i <= j; i, j = i+1, j-1 {
       nums[i], nums[j] = nums[j], nums[i]
   }
   return nums 
}
```

