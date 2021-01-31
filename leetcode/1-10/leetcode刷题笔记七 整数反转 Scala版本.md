leetcode刷题笔记七 整数反转 Scala版本

源地址：(leetcode刷题笔记七 整数反转 Scala版本)[https://leetcode.com/problems/reverse-integer/submissions/]

问题描述：

>Given a 32-bit signed integer, reverse digits of an integer.

> **Example 1:**

```
Input: 123
Output: 321
```

> **Example 2:**

```
Input: -123
Output: -321
```

> **Example 3:**

```
Input: 120
Output: 21
```

> **Note:**

> Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

简要思路分析：

> 对于整数反转，可以从字符串或者Long型数据处理。对于字符串处理方法，我们需要将符号与数字剥离，通过对原符号推出转换后的符号，对气候的数字串进行反转，将首部的0处理掉；对于Long型数据处理，针对X不断取余和除10 将个位数分离出来作为转换后的数的高位处理。这两种都要解决的问题是Int型数反转后溢出的情况进行处理，这是本题的关键！

代码补充：

> 法一：
>
> ```scala
> object Solution {
> def reverse(x: Int): Int = {
>         var xStr = x.toString
>         var Xbol = ""
>         if (xStr.length == 0 || xStr.length == 1) return xStr.toInt
>         if (xStr.charAt(0) == '-') {
>           Xbol = "-"
>           xStr = xStr.substring(1).reverse
>         }
>         else xStr = xStr.reverse
>         while(xStr.charAt(0) == '0'){
>           xStr = xStr.substring(1)
>         }
>         //println(Xbol.concat(xStr))
>         val x1 = (Xbol.concat(xStr)).toLong
>         if (x1 > Int.MaxValue || x1 < Int.MinValue) {
>             return 0
>         }
>         else
>         return x1.toInt
>     }
> }
> ```
>
> 法二：
>
> ```scala
> object Solution {
> def reverse(x: Int): Int = {
> 
>       if (x > Int.MaxValue || x < Int.MinValue) {
>         return 0
>       }
> 
>       var xTemp = x
>       var result:Long = 0
>       while (xTemp != 0){
>         result = result * 10 + xTemp % 10
>         xTemp /= 10
>       }
> 
>       if (result > Int.MaxValue || result < Int.MinValue) {
>         return 0
>       }
>       else
>         return  result.toInt
> 
>     }
> }
> ```
>
> 