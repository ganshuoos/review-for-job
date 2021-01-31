6. Z 字形变换

源地址：[6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/)

问题描述：

>将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
>
>比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
>
>P   A   H   N
>A P L S I I G
>Y   I   R
>之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
>
>请你实现这个将字符串进行指定行数变换的函数：
>
>string convert(string s, int numRows);
>
>
>示例 1：
>
>输入：s = "PAYPALISHIRING", numRows = 3
>输出："PAHNAPLSIIGYIR"
>示例 2：
>输入：s = "PAYPALISHIRING", numRows = 4
>输出："PINALSIGYAHRPI"
>解释：
>P     I    N
>A   L S  I G
>Y A   H R
>P     I
>示例 3：
>
>输入：s = "A", numRows = 1
>输出："A"
>
>
>提示：
>
>1 <= s.length <= 1000
>s 由英文字母（小写和大写）、',' 和 '.' 组成
>1 <= numRows <= 1000

``` go
func convert(s string, n int) string {
    length := len(s)

    if n == 1 {return s}
    res := make([]byte, 0)

    for i := 0; i < n; i++ {
        if i == 0 || i == n-1 {
            for j := i; j < length; j += 2*n-2 {
                res = append(res, s[j])
            }
        } else  {
            for j, k := i, 2*n-i-2; j < length || k < length; j, k = j + 2*n-2, k+2*n-2 {
                if j < length {
                    res = append(res, s[j])
                }
                if k < length {
                    res = append(res, s[k])
                }
            }   
        }
    }

    return string(res)
}
```



