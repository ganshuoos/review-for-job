剑指 Offer 09. 用两个栈实现队列

地址：[剑指 Offer 09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

问题描述：

>用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
>
> 
>
>示例 1：
>
>输入：
>["CQueue","appendTail","deleteHead","deleteHead"]
>[[],[3],[],[]]
>输出：[null,null,3,-1]
>示例 2：
>
>输入：
>["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
>[[],[],[5],[2],[],[]]
>输出：[null,-1,null,null,5,2]
>提示：
>
>1 <= values <= 10000
>最多会对 appendTail、deleteHead 进行 10000 次调用

``` scala
import scala.collection.mutable.Stack

class CQueue() {
    val fStack = Stack[Int]()
    val sStack = Stack[Int]()

    def appendTail(value: Int) {
        fStack.push(value)
    }

    def deleteHead(): Int = {
        if (sStack.isEmpty == false) {
            return sStack.pop
        } else {
            while (fStack.size > 0) {
                sStack.push(fStack.pop)
            }
            if (sStack.isEmpty == true) {
                return -1;
            } else {
                return sStack.pop
            }
        }
    }

}

/**
 * Your CQueue object will be instantiated and called as such:
 * var obj = new CQueue()
 * obj.appendTail(value)
 * var param_2 = obj.deleteHead()
 */
```

```go
type CQueue struct {
    inStack Stack
    outStack Stack 
}

type Stack []int

func (s *Stack) Push (value int) {
    *s = append(*s, value)
}

func (s *Stack) Pop() int {
    n := len(*s)
    res := (*s)[n-1]

    *s =  (*s)[0:n-1]
    return res
}

func Constructor() CQueue {
    return CQueue{}
}



func (this *CQueue) AppendTail(value int)  {
    this.inStack = append(this.inStack, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.outStack) != 0 {
        return this.outStack.Pop()
    } else if len(this.inStack) != 0 {
        for (len(this.inStack) != 0){
            this.outStack.Push(this.inStack.Pop())
        }
        return this.outStack.Pop()
    }
    return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
```

