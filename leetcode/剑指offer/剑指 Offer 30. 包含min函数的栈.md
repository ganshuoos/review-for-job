剑指 Offer 30. 包含min函数的栈

地址：[剑指 Offer 30. 包含min函数的栈](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/)

>定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
>
> 
>
>示例:
>
>MinStack minStack = new MinStack();
>minStack.push(-2);
>minStack.push(0);
>minStack.push(-3);
>minStack.min();   --> 返回 -3.
>minStack.pop();
>minStack.top();      --> 返回 0.
>minStack.min();   --> 返回 -2.
>
>
>提示：
>
>各函数的调用总次数不超过 20000 次
>

``` scala

```

```go
type MinStack struct {
    Stack   []int
    TStack  []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        Stack   : []int{},
        TStack    : []int{},
    }
}


func (this *MinStack) Push(x int)  {
    this.Stack = append(this.Stack, x)
    if len(this.TStack) == 0 || this.TStack[len(this.TStack)-1] >= x {
        this.TStack = append(this.TStack, x)
    }
}


func (this *MinStack) Pop()  {
    popValue := this.Stack[len(this.Stack)-1]
    if this.TStack[len(this.TStack)-1] >= popValue {
        this.TStack = this.TStack[0:len(this.TStack)-1]
    }
    this.Stack = this.Stack[0:len(this.Stack)-1]
}


func (this *MinStack) Top() int {
    //fmt.Println(this.Stack)
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) Min() int {
    //fmt.Println(this.TStack)
    return this.TStack[len(this.TStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
```

