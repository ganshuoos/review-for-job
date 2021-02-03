28. 实现 strStr()

源地址：[28. 实现 strStr()](https://leetcode-cn.com/problems/implement-strstr/)

问题描述：

>实现 strStr() 函数。
>
>给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
>
>示例 1:
>
>输入: haystack = "hello", needle = "ll"
>输出: 2
>示例 2:
>
>输入: haystack = "aaaaa", needle = "bba"
>输出: -1
>说明:
>
>当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
>
>对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
>

``` go
func strStr(haystack string, needle string) int {
    if len(needle) == 0  {return 0}
    if len(haystack) < len(needle) {return -1}

    for i := 0; i < len(haystack) - len(needle) + 1; i++ {
        if string(haystack[i: i+len(needle)]) == needle {
            return i
        } 
    }
    return -1
}

//KMP 算法
func strStr(haystack string, needle string) int {
    if len(needle) == 0 {return 0}
    haystack = " " + haystack
    needle = " " + needle

    next := make([]int, len(needle))

    for i, j := 2, 0; i < len(needle); i++ {
        for j > 0 && needle[i] != needle[j+1] {j = next[j]}
        if needle[i] == needle[j+1] {j++}
        next[i] = j
    }


    for i, j := 1, 0; i < len(haystack); i++ {
        for j > 0 && haystack[i] != needle[j+1] {j = next[j]}
        if haystack[i] == needle[j+1] {j++}
        if j == len(needle)-1 {
            return i - len(needle)+1
        }
    }

    return -1 
}
```



