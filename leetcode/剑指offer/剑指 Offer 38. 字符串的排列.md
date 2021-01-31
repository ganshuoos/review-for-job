剑指 Offer 38. 字符串的排列

地址：[剑指 Offer 38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)

>输入一个字符串，打印出该字符串中字符的所有排列。
>
> 
>
>你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
>
> 
>
>示例:
>
>输入：s = "abc"
>输出：["abc","acb","bac","bca","cab","cba"]
>
>
>限制：
>
>1 <= s 的长度 <= 8

``` scala

```

```go
//var res []string
//var path []byte
var st []bool
var judgeMap map[string]bool

func permutation(s string) []string {
    length := len(s)
    path := make([]byte, 0)
    res := make([]string, 0)
    judgeMap = make(map[string]bool)
    st = make([]bool, length)
    dfs(s, 0, path, &res)
    return res
}

func dfs(s string, pathLen int, path []byte, res *[]string) {
    if pathLen == len(s) {
        //fmt.Println(path)
        if _ ,ok := judgeMap[string(path)]; !ok {
            (*res) = append(*res, string(path))
            judgeMap[string(path)] = true
        }

    }
    for i := 0; i < len(s); i++ {
        if st[i] == true {
            continue
        }
        st[i] = true
        path = append(path, s[i])
        dfs(s, pathLen+1, path, res)
        path = path[:len(path)-1]
        st[i] = false
    }
}
```

