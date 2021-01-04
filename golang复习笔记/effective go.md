**effctive go读书笔记**

**1.短变量声明**

>参考博客：
>
>https://zhuanlan.zhihu.com/p/259182614
>
>https://blog.csdn.net/wfyeshi/article/details/60572739
>
>https://blog.csdn.net/u013600225/article/details/78628959
>
>满足以下条件时，已声明变量可出现在短变量声明中：
>
>1. 两次声明处于同一作用域内
>2. 其类型相应
>3. 在此次声明中至少另有一个变量为新声明的
>
>```go
>//举例而言
>package main
>
>import (
>	"errors"
>	"fmt"
>)
>
>func main()  {
>	f, err := 1, errors.New("error1")
>	if err != nil {
>		_ = f
>		fmt.Println(&err)
>        //0xc0000361f0
>	}
>
>    //这里err创建了新的变量，局部变量覆盖全局变量err
>	if d, err := 2, errors.New("error2"); err != nil {
>		_ = d
>		fmt.Println(&err)
>		//0xc000036210
>	}
>
>    //这里相同的作用域下，err变量并没有重新创建，而是直接进行赋值
>	a, err := 3, errors.New("error3")
>	fmt.Println(a, err)
>    //3 error3
>	fmt.Println(&err)
>    //0xc0000361f0
>}
>
>
>//注意由短变量声明导致的以下错误
>func() (err error) {  
>        aa, err := 1, errors.New(" a error") // a是新变量，err是被赋值  
>        if err != nil {  
>            return // 正确返回err  
>        }  
>              //  ------------------------------------------------  
>        if bb, err := 2,  errors.New("b error"); err != nil { // 此刻if语句中err被重新创建  
>            return      // if语句中的err覆盖外面的err，导致编译  错误 (err is shadowed during return) 
>         //无返回值的return,会把返回值列表中的err,而不是新的作用域里的err
>        //解决方案----------------------------  
>    	//if bb, err1 := 2,  errors.New("b error"); err1 != nil {  
>    	//  err = err1  
>    	//  return  
>    	//-------------------------------------  
>     	} else {   
>        	fmt.Println(bb)   
>    	}   
>    return   
>}()
>
>
>package main
>
>import "fmt"
>
>func main() {
>	myTest()
>	v0, v1 := myTest3()
>	fmt.Println("myTest3:", v0, v1)
>}
>
>func myTest() (retI int, retF float64) { //返回值列表定义了retI和retF变量,作用域是整个函数体
>	//retI, retF := 0, 0.0 //报错:[no new variables on left side of :=]
>	//左边不需要"全部都是新变量",但可以是"一部分新变量+一部分老变量"
>	i0, retF := 1, 1.1
>	i1, retF := 2, 2.2
>	i2, retF := 3, 3.3
>	fmt.Println("print1:", retI, retF, i0, i1, i2)
>
>	if true { //新的语句块,产生了新的作用域
>		retI, retF := -1, -1.1 //这里又定义了新的变量retI和retF,和返回值列表重名了,作用域是if语句块
>		fmt.Println("print2:", retI, retF)
>	}
>
>	fmt.Println("print3:", retI, retF)
>
>	if true {
>		j0, retF := -2, -2.2 //这里又定义了新的变量j0和retF,其中retF和返回值列表中的retF重名了
>		fmt.Println("print4:", retI, retF, j0)
>	}
>
>	fmt.Println("print5:", retI, retF)
>
>	if true {
>		retI, retF = 9, 9.9 //修改老变量的值
>		fmt.Println("print6:", retI, retF)
>	}
>
>	fmt.Println("print7:", retI, retF)
>
>	return
>}
>
>func myTest2() (retI int, retF float64) {
>	if true { //新的作用域
>		i0, retF := -1, -1.1
>		fmt.Println(i0, retF)
>		return //报错:[retF is shadowed during return]
>		//无返回值的return,会把返回值列表中的retI和retF返回出去,而不是新的作用域里的retF
>	}
>	return
>}
>
>func myTest3() (retI int, retF float64) {
>	if true {
>		i0, retF := -1, -1.1
>		fmt.Println(i0, retF)
>		return retI, retF //这里返回的是外层的retI和里层的retF,因为指明了,所以没有报错
>	}
>	return
>}
>```

**2.函数**

> 函数多返回值
>
> (1)便于编写代码 （2）关注error的处理
>
>  
>
> 浅谈defer
>
> （1）defer队列先进后出 （2）设置返回值  defer 执行RET指令
>
> 