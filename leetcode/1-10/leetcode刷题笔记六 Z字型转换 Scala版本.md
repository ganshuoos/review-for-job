leetcode刷题笔记六 Z字型转换 Scala版本

源地址：(leetcode刷题笔记六 Z字型转换 Scala版本)[https://leetcode.com/problems/zigzag-conversion/submissions/]

问题描述：

>The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

> And then read line by line: `"PAHNAPLSIIGYIR"`

> Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);
```

> **Example 1:**

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

> **Example 2:**

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```

简要思路分析：

> 设置层数指示器与方向参数，通过判断执行是否达到顶部或者底部，进行方向变更，将对应层的字符放入对应的字符串数组中，最后将字符串数组进行汇总。

代码补充：

```scala
object Solution {
    def convert(s: String, numRows: Int): String = {
         if (numRows == 1 || s.length <= numRows) return s

      var upDown = true
      var counter = 0
      var arr:Array[String] = new Array[String](numRows)
      var answer = ""


      for (i <- 0 until s.length){
        //println("--------------------------")
        //println("i: " + i.toString)
        //println("counterPre: "+counter.toString)
        //println("upDownPre: "+upDown.toString)
        arr(counter) += s.charAt(i)
        /* 用match不可以。。 神秘
        (counter) match{
          case(0) => upDown = true
          case(maxCounter) => upDown = false
        }
        */
        if (counter == 0) upDown = true
        if (counter == numRows-1) upDown = false
        (upDown) match{
          case(true) => counter+=1
          case(false) => counter-=1
        }
        //println("counterAfter: "+counter.toString)
        //println("upDownAfter: "+upDown.toString)
        //println("s(i): " + s.charAt(i))
      }

      for (i <- 0 until numRows){
        answer += arr(i).drop(4).toString
        //println(arr(i))
    }
     return answer
    }
}
```

