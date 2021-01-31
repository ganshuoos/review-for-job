剑指 Offer 05. 替换空格

地址：[剑指 Offer 05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

问题描述：

>请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
>
> 
>
>示例 1：
>
>输入：s = "We are happy."
>输出："We%20are%20happy."
>
>
>限制：
>
>0 <= s 的长度 <= 10000
>

``` scala
//也可以考虑replace函数
object Solution {
    def replaceSpace(s: String): String = {
        var res = ""
        for (i <- 0 to s.length-1) {
            if (s(i) == ' ') {
                res += "%20"
            }  else {
                res += s(i)
            }
        }
        return res
    }
}
```

```go
func replaceSpace(s string) string {
    res := ""
    for _, str := range s {
        if str == ' ' {
            res += "%20"
        } else {
            res += string(str)
        }
    }
    return res 
}
```

