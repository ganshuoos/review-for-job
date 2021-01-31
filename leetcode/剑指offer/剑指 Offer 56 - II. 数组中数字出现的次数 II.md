剑指 Offer 56 - II. 数组中数字出现的次数 II

地址：[剑指 Offer 56 - II. 数组中数字出现的次数 II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)

> 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
>
>  
>
> 示例 1：
>
> 输入：nums = [3,4,3,3]
> 输出：4
> 示例 2：
>
> 输入：nums = [9,1,7,9,7,9,7]
> 输出：1
>
>
> 限制：
>
> 1 <= nums.length <= 10000
> 1 <= nums[i] < 2^31

``` 

```

```go
import "fmt"
func singleNumber(nums []int) int {
    one, two := 0, 0
    for _, value := range nums {
        //go语言中取反使用＾　如果作为二元运算符，^ 表示按位异或，即：对应位相同为 0，相异为 1。操作符 &^，按位置零，例如：z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。

        one = (one ^ value) & (^two)
        two = (two ^ value) & (^one)
    }
    //three := 2
    //fmt.Println(fal(three))
    return one
}
```

