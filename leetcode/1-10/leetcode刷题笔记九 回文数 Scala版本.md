leetcode刷题笔记九 回文数 Scala版本

源地址：[leetcode刷题笔记九 回文数 Scala版本](https://leetcode.com/problems/palindrome-number/)

问题描述：

>Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

> **Example 1:**

```
Input: 121
Output: true
```

> **Example 2:**

```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

> **Example 3:**

```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

简要思路分析：

> 判断回文的主要方式是比较原串与反转串是否一致或者串本身是否具有对称性。由于Scala具有相关函数，这里我们简单比较一下原串与反转串是否一致即可。

代码补充：

```scala
object Solution {
    def isPalindrome(x: Int): Boolean = {
      if (x.toString == x.toString.reverse) return true  else return false
    }
}
```



