剑指 Offer 58 - I. 翻转单词顺序

地址：[剑指 Offer 58 - I. 翻转单词顺序](https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/)

> 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
>
>  
>
> 示例 1：
>
> 输入: "the sky is blue"
> 输出: "blue is sky the"
> 示例 2：
>
> 输入: "  hello world!  "
> 输出: "world! hello"
> 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
> 示例 3：
>
> 输入: "a good   example"
> 输出: "example good a"
> 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
>
>
> 说明：
>
> 无空格字符构成一个单词。
> 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
> 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

``` 

```

```go
import (
    "fmt"
    "strings"
)
func reverseWords(s string) string {
    //使用strings.Trim去除首尾的kongge
    s = strings.Trim(s, " ")
    ans := ""
    
    //从尾部开始扫描
    i, j := len(s)-1, len(s)-1

    for i >= 0 {
        //从非空格末尾开始向前遍历， 遇到第一个空格（即完整单词）将其取出
        for i >= 0 && s[i] != ' ' {
            i -= 1
        }
        //添加进入ans
        ans += s[i+1:j+1]
        ans += " "
        //去除单词间的空格
        for i >= 0 && s[i] == ' ' {
            i -= 1
        }
        //移动j指针
        j = i
    }
    return strings.Trim(ans, " ")
}
```

