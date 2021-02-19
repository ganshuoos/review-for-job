93. 复原IP地址

源地址：[93. 复原IP地址](https://leetcode-cn.com/problems/restore-ip-addresses/)

问题描述：

>给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
>
>有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
>
>例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
>
> 
>
>示例 1：
>
>输入：s = "25525511135"
>输出：["255.255.11.135","255.255.111.35"]
>示例 2：
>
>输入：s = "0000"
>输出：["0.0.0.0"]
>示例 3：
>
>输入：s = "1111"
>输出：["1.1.1.1"]
>示例 4：
>
>输入：s = "010010"
>输出：["0.10.0.10","0.100.1.0"]
>示例 5：
>
>输入：s = "101023"
>输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
>
>
>提示：
>
>0 <= s.length <= 3000
>s 仅由数字组成

``` go
var res []string
var path string

func restoreIpAddresses(s string) []string {
    //path = ""
    res = make([]string, 0)
    dfs(s, 0, 0, "")
    return res
}

func dfs(s string, u int, k int, path string) {
    if u == len(s) {
        if k == 4 {
            path = path[:len(path)-1]
            fmt.Println(path)
            tmp := make([]byte, len(path))
            copy(tmp, path)
            res = append(res, string(tmp))
        }
        return
    }

    if k == 4 {return}

    for  i, t := u,  0; i < len(s); i++ {
        if i > u && s[u] == '0' {break}
        t = t * 10 + int(s[i] - '0')

        if t <= 255 {
            newPath := path + strconv.Itoa(t) + "."
            dfs(s, i+1, k+1, newPath)
        } else {
            break
        }
    }
}
```



