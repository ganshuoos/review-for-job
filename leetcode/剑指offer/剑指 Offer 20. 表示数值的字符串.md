剑指 Offer 20. 表示数值的字符串

地址：[剑指 Offer 20. 表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

问题描述：

>请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
>

``` scala
//思路参考剑指offer解法及题解https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/zui-jian-dan-si-lu-xiang-xi-zhu-shi-zheng-shu-xiao/
//主要构建有符号整数和无符号整数判断两个辅助函数
//根据 . 和 e 两侧数的要求结合判断
object Solution {
    var index = 0
    def isNumber(s: String): Boolean = {
        if (s == null || s.length == 0) {return false}
        index = 0
        val str = s + "|"
        while (str(index) == ' ') {index += 1}
        var numberic = scanInteger(str)

        if (str(index) == '.') {
            index += 1
            numberic = scanUnsignedInteger(str) || numberic
        }

        if (str(index) == 'e' || str(index) == 'E') {
            index += 1
            numberic = scanInteger(str) && numberic
        }

        while (str(index) == ' ') {index += 1}

        return numberic && str(index) == '|'
    }

    def scanInteger(str: String): Boolean = {
        if (str(index) == '+' || str(index) == '-') {index += 1}
        return scanUnsignedInteger(str)
    }

    def scanUnsignedInteger(str: String): Boolean = {
        val before = index
        while (str(index) >= '0' && str(index) <= '9') {
            index += 1
        }
        return index > before
    }
}
```

```go
import "fmt"

var index int

func isNumber(s string) bool {
    index = 0
    if len(s) == 0  { return false}

    s += "|"

    for s[index] == ' ' { index += 1}
    numberic := scanInteger(s)
    //fmt.Println(numberic)

    if s[index]  == '.' {
        index += 1
        numberic = scanUnsignedInteger(s) || numberic
    }

    if s[index] == 'E' || s[index] == 'e' {
        index += 1
        numberic = scanInteger(s) && numberic
    }
    fmt.Println(numberic)
    for s[index] == ' ' { index += 1}
    return numberic && s[index] == '|'
}

func scanUnsignedInteger(str string) bool {
    before := index
    for str[index] >= '0' && str[index] <= '9'  {index += 1}
    return index > before
}

func scanInteger(str string) bool {
    if str[index] == '+' || str[index] == '-' {index += 1}
    return scanUnsignedInteger(str)
}
```

