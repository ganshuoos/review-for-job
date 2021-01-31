leetcode刷题笔记八 字符串转整性 Scala版本

源地址： [leetcode刷题笔记八 字符串转整性 Scala版本](https://www.cnblogs.com/ganshuoos/p/12663458.html)

问题描述：

>

- Only the space character `' '` is considered as whitespace character.
- Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

> **Example 1:**

```
Input: "42"
Output: 42
```

> **Example 2:**

```
Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
```

> **Example 3:**

```
Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
```

> **Example 4:**

```
Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical 
             digit or a +/- sign. Therefore no valid conversion could be performed.
```

> **Example 5:**

```
Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
```

简要思路分析：

> 本题的关键点主要是判断AtoI的极端形式，主要是筛选出失败的字符串形式。一种是正常对Str进行处理，严格按照各种形式去判断Str是否合法；另一种是通过正则表达式完成对Str的筛选，然后对Str进行分析判断，这里我才用的正则方法。

代码补充：

>```scala
>object Solution {
>   def myAtoi(str: String): Int = {
>      """^\s*([+-]?)0*(\d+)""".r.findFirstMatchIn(str) match {
>
>          case Some(m) => {
>          val sign = m.group(1)
>            println(sign)
>          val nums = m.group(2)
>            println(nums)
>          (sign+nums).toDouble match {
>            case n if n > Int.MaxValue => return Int.MaxValue
>            case n if n < Int.MinValue => return Int.MinValue
>            case n => return n.toInt
>          }
>          }
>          case None => return 0
>      }
>    }
>}
>```

知识补充：

> Scala正则相关知识：
>
> | 表达式    | 匹配规则                                                     |
> | :-------- | :----------------------------------------------------------- |
> | ^         | 匹配输入字符串开始的位置。                                   |
> | $         | 匹配输入字符串结尾的位置。                                   |
> | .         | 匹配除"\r\n"之外的任何单个字符。                             |
> | [...]     | 字符集。匹配包含的任一字符。例如，"[abc]"匹配"plain"中的"a"。 |
> | [^...]    | 反向字符集。匹配未包含的任何字符。例如，"[^abc]"匹配"plain"中"p"，"l"，"i"，"n"。 |
> | \\A       | 匹配输入字符串开始的位置（无多行支持）                       |
> | \\z       | 字符串结尾(类似$，但不受处理多行选项的影响)                  |
> | \\Z       | 字符串结尾或行尾(不受处理多行选项的影响)                     |
> | re*       | 重复零次或更多次                                             |
> | re+       | 重复一次或更多次                                             |
> | re?       | 重复零次或一次                                               |
> | re{ n}    | 重复n次                                                      |
> | re{ n,}   |                                                              |
> | re{ n, m} | 重复n到m次                                                   |
> | a\|b      | 匹配 a 或者 b                                                |
> | (re)      | 匹配 re,并捕获文本到自动命名的组里                           |
> | (?: re)   | 匹配 re,不捕获匹配的文本，也不给此分组分配组号               |
> | (?> re)   | 贪婪子表达式                                                 |
> | \\w       | 匹配字母或数字或下划线或汉字                                 |
> | \\W       | 匹配任意不是字母，数字，下划线，汉字的字符                   |
> | \\s       | 匹配任意的空白符,相等于 [\t\n\r\f]                           |
> | \\S       | 匹配任意不是空白符的字符                                     |
> | \\d       | 匹配数字，类似 [0-9]                                         |
> | \\D       | 匹配任意非数字的字符                                         |
> | \\G       | 当前搜索的开头                                               |
> | \\n       | 换行符                                                       |
> | \\b       | 通常是单词分界位置，但如果在字符类里使用代表退格             |
> | \\B       | 匹配不是单词开头或结束的位置                                 |
> | \\t       | 制表符                                                       |
> | \\Q       | 开始引号：**\Q(a+b)\*3\E** 可匹配文本 "(a+b)*3"。            |
> | \\E       | 结束引号：**\Q(a+b)\*3\E** 可匹配文本 "(a+b)*3"。            |
>
> 正则表达式实例
>
> | 实例            | 描述                                          |
> | :-------------- | :-------------------------------------------- |
> | .               | 匹配除"\r\n"之外的任何单个字符。              |
> | [Rr]uby         | 匹配 "Ruby" 或 "ruby"                         |
> | rub[ye]         | 匹配 "ruby" 或 "rube"                         |
> | [aeiou]         | 匹配小写字母 ：aeiou                          |
> | [0-9]           | 匹配任何数字，类似 [0123456789]               |
> | [a-z]           | 匹配任何 ASCII 小写字母                       |
> | [A-Z]           | 匹配任何 ASCII 大写字母                       |
> | [a-zA-Z0-9]     | 匹配数字，大小写字母                          |
> | [^aeiou]        | 匹配除了 aeiou 其他字符                       |
> | [^0-9]          | 匹配除了数字的其他字符                        |
> | \\d             | 匹配数字，类似: [0-9]                         |
> | \\D             | 匹配非数字，类似: [^0-9]                      |
> | \\s             | 匹配空格，类似: [ \t\r\n\f]                   |
> | \\S             | 匹配非空格，类似: [^ \t\r\n\f]                |
> | \\w             | 匹配字母，数字，下划线，类似: [A-Za-z0-9_]    |
> | \\W             | 匹配非字母，数字，下划线，类似: [^A-Za-z0-9_] |
> | ruby?           | 匹配 "rub" 或 "ruby": y 是可选的              |
> | ruby*           | 匹配 "rub" 加上 0 个或多个的 y。              |
> | ruby+           | 匹配 "rub" 加上 1 个或多个的 y。              |
> | \\d{3}          | 刚好匹配 3 个数字。                           |
> | \\d{3,}         | 匹配 3 个或多个数字。                         |
> | \\d{3,5}        | 匹配 3 个、4 个或 5 个数字。                  |
> | \\D\\d+         | 无分组： + 重复 \d                            |
> | (\\D\\d)+/      | 分组： + 重复 \D\d 对                         |
> | ([Rr]uby(, )?)+ | 匹配 "Ruby"、"Ruby, ruby, ruby"，等等         |
>
> 正则表达式捕获与非捕获
>
> 这一部分主要参考https://blog.csdn.net/liuxiao723846/article/details/83277488 和  https://blog.csdn.net/liuxiao723846/article/details/83278067
>
> 分组：正则表达式中的**分组**又称为**子表达式**，就是把一个正则表达式的全部或部分当做一个整体进行处理，分成一个或多个组。其中分组是使用“（）”表示的。进行分组之后“（）”里面的内容就会被当成一个整体来处理。
>
> 分组进一步可分为捕获组与非捕获组。
>
> 1、捕获组（capture group）：
>
> 捕获组：就是把正则表达式中子表达式匹配的内容，保存到内存中以数字编号或显式命名的组里，方便后面引用。当然，这种引用既可以是在正则表达式内部，也可以是在正则表达式外部。
>
> 捕获组编号规则：
>
> 从左向右，以分组的左括号为标志，第一个出现的分组的组号为1，第二个为2，以此类推。需要注意的是组0永远代表的是整个正则式
>
> 2、非捕获组
>
> 1）为什么要有非捕获组：
>
> 一旦使用了“()”，就会默认为是捕获组，从而将“()”内表达式匹配的内容捕获到组里。但是有些情况下，不得不用“()”，但并不关心“()”中匹配的内容是什么，后面也不会引用捕获到的内容，这带来了一个副作用，记录这些捕获组就会占用内存，降低匹配效率。设计非捕获组的目的就是为了抵消这种副作用。 只进行分组，并不将子表达式匹配到的内容捕获到组里。
>
> 2）使用：
>
> 以 (?) 开头的组是非捕获组，它不捕获文本 也不针对组合计进行计数。就是说，如果小括号中以?号开头，那么这个分组就不会捕获文本，当然也不会有组的编号，因此也不存在反向引用。
>
> 当一个正则表达式被分组后，每个组将会自动的分配一个组号用于代表该组的表达式，其中，组号的编制规则为：从左到右、以分组的左括号“(”为标志，第一个分组的组号为1，第二个分组的组号为2，以此类推。反向引用提供查找重复字符组的方便的方法。它们可被认为是再次匹配同一个字符串的**快捷指令**。
>
> 如本文中使用1组捕获sign标志，2组捕获nums标志
>
> ```
> (boy)\1          //相当于(boy)(boy),匹配boyboy
> (boy)(girl)\1\2  //匹配boygirlboygirl
> ===============================================
> 对比(\w)\1和的区别：
> (\w)\1    //该正则匹配的是出现两次的字符,\1将重复匹配(\w）出现的内容！例如匹配aa、bb
> (\w)(\w)  //该正则只是简单的匹配两个字符，例如匹配ab，aa等。注意区别！
> ```
>
> 