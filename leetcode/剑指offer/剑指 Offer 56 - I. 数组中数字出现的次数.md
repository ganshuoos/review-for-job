剑指 Offer 56 - I. 数组中数字出现的次数

地址：[剑指 Offer 56 - I. 数组中数字出现的次数](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)

> 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
>
>  
>
> 示例 1：
>
> 输入：nums = [4,1,4,6]
> 输出：[1,6] 或 [6,1]
> 示例 2：
>
> 输入：nums = [1,2,10,4,1,4,3,3]
> 输出：[2,10] 或 [10,2]
>
>
> 限制：
>
> 2 <= nums.length <= 10000

``` 

```

```go
func singleNumbers(nums []int) []int {
    res := make([]int, 0)

    //由于两个不相同的数， 异或结果不为0 即某一位肯定为1
    //这情况，这个位置 a与其取与等于1， b等于0 
    num := 0
    for _, value := range nums {
        num ^= value
    }

    div := 1
    for div&(num) == 0 {
        div = div << 1
    }

    
    second := 0
    for _, value := range nums {
        if (value & div) == 1 {
            res = append(res, value)
        } else {
            //利用其他的两个异或全部去除 只剩最后一个目标数
            second ^= value  
        }    
    }
    res = append(res, second)
    return res
}
```

