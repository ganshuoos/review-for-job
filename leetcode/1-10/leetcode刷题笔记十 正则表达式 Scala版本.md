leetcode刷题笔记十 正则表达式 Scala版本

源地址：[leetcode刷题笔记十 正则表达式 Scala版本](https://leetcode.com/problems/regular-expression-matching/submissions/)

问题描述：

>Given an input string (`s`) and a pattern (`p`), implement regular expression matching with support for `'.'` and `'*'`.

```
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

> The matching should cover the **entire** input string (not partial).

**Note:**

- `s` could be empty and contains only lowercase letters `a-z`.
- `p` could be empty and contains only lowercase letters `a-z`, and characters like `.` or `*`.

**Example 1:**

```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**

```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**

```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Example 4:**

```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

**Example 5:**

```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```

简要思路分析：

> 首先对正则表达式进行分析，对于s串，其对应小写字母a-z；对于p串，其对应小写字母a-z，.或者*。因为对于样式，每个字符有三种可能，我们对三种可能进行分析。
>
> s[i] : 字符串s第i个字符的位置  p[j]:样式串p第j个字符的位置
>
> match[i]:为第i个状态的正则匹配情况
>
> （1）p[j] 为一个字母
>
> 设 s="aa" p="ab" 
>
> s[0] == p[0] 	==> match[0] = T 
>
> s[1] != p[1]  ==> match[1] = (match[0] && (s[1] == p[1])) => F 
>
> （2）p[j] = '.',这里'.'代表'.' 匹配任意单个字符
>
> 设s="aa", p="a." 
>
> 由于 '.'可以匹配任意单个字符，当p[j]='.',我们不需要考虑对应的s串的字符是什么。
>
> s[0] == p[0] 	==> match[0] = T 
>
> s[1] != p[1] 但 p[1] ='.'  ==> match[1] = (match[0] && (p[1] == '.')) => T
>
> (3) p[j] = '*'， *匹配零个或多个前面的那一个元素
>
> 因为'*'存在两种情况，0次匹配或者多次匹配
>
> 就0次匹配而言，视其对于'*'号前一字符的影响也存在两种情况。
>
> 【1】0次匹配，且对前一字符消失
>
> 设s="a", p="ac*" 这里p[2]='\*',且*'*'令前一个字符C消失，这里match[0]\[2\] = match[0]\[0\]  j向前移动两位
>
> 【2】 0次匹配，不影响前一字符
>
> 设s="a",p="ac\*" 或者 s = "a" , p=".\*", 这里p[1]="\*" 
>
> match[0]\[2\] = match[0]\[0\]  ==> if ((s[0] == p[0]) || p[0] == '.') 
>
>  【3】多次匹配，不影响前一字符
>
> 设s="aa",p="a*" ==> match[1]\[1\] = match[1]\[0\] 
>
> 总结状态方程
>
> match[i]\[j\] = match[i-1]\[j-1\]  		if p[j-1] != '*'  && ((s[i-1] == p[j-1]) || p[j-1] == '.' )
>
> match\[i\][j] = match\[i\][j-2]		if  p[j-1] == '\*'  && （s[i-1] != p[j-2]）
>
> match\[i\][j] = match[i-1]\[j\] || match\[i\][j-2]  if p[j-1] == '\*' && ((s[i-1] != p[j-2]) || p[j-2] == '.') 
>
> 根据状态方程，可以通过回溯法或者动态规划实现

代码补充：

> 1.回溯法
>
> ```scala
> object Solution {
>     def isMatch(s: String, p: String): Boolean = {
>       if (p.length <= 0 ) return s.length <= 0
>       var imatch:Boolean = (s.length > 0 &&(s.charAt(0) == p.charAt(0) || p.charAt(0) == '.'))
>       //println("imatch: " + imatch)
>       if (p.length > 1 && p.charAt(1) == '*'){
>         return isMatch(s, p.substring(2)) ||( imatch && isMatch(s.substring(1), p))
>       }
>       else{
>         return imatch && (isMatch(s.substring(1),p.substring(1)))
>       }
> 
>     }
> }
> ```
>
> 2.动态规划 
>
> ```scala
> object Solution {
>     def isMatch(s: String, p: String): Boolean = {
>       var dp = Array.ofDim[Boolean](s.length+2, p.length+2)
>       //var dp = Array.fill(s.length, p.length)(false)
>       val sStr = ' ' + s
>       val pStr = ' ' + p
>       if (pStr.length <= 0) return sStr.length<=0
>       dp(0)(0) = true
> 
>       for (i <- 1 to sStr.length){
>         for(j <- 1 to pStr.length){
>           //println("----------------------------")
>           //println("i: "+i)
>           //println("j: "+j)
>           //println("p(j-1): " + pStr.charAt(j-1))
>           //println("s(i-1): " + sStr.charAt(i-1))
> 
>           if (pStr.charAt(j-1) != '*'){
>             if (pStr.charAt(j-1) == sStr.charAt(i-1) || pStr.charAt(j-1) == '.'){dp(i)(j) = dp(i-1)(j-1)}
>             else {dp(i)(j)=false}
>             //println("!*")
>             //println("dp(i-1)(j-1): "+dp(i-1)(j-1))
>           }
>           else{
>             if (pStr.charAt(j-2) == sStr.charAt(i-1) || pStr.charAt(j-2) == '.') dp(i)(j) = dp(i)(j-2) || dp(i-1)(j)
>             else dp(i)(j) = dp(i)(j-2)
>             //dp(i)(j) = dp(i)(j-2)||dp(i-1)(j)
>             //println("*")
>             //println("dp(i)(j-2): "+dp(i)(j-2))
>             //println("dp(i-1)(j): "+dp(i-1)(j))
>           }
>           //println("dp(i)(j): " + dp(i)(j))
>           //println("dp(i)(j): " + dp(i)(j))
>         }
>       }
>       return dp(s.length+1)(p.length+1)
>     }
> }
> ```
>
> 