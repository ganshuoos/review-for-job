**go 调试工具 – delve**

> 参考:https://www.cnblogs.com/wanghui-garcia/p/10455217.html
>
> 类似于GDB工具， 可以查看一些具有go特性的调试工具。
>
> **安装**
>
> 由于网络环境问题，需要先进行网络设置。
>
> ```go
> set GOPROXY=https://goproxy.cn
> set GO111MODULE=on
> //获取delve
> go get github.com/go-delve/delve/cmd/dlv
> ```
>
> **delve debug**
>
> ```go
> /*
> The following commands are available:
>     args ------------------------ 打印函数参数.
>     break (alias: b) ------------ 设置断点.
>     breakpoints (alias: bp) ----- 输出活动断点的信息.
>     call ------------------------ 恢复进程，注入一个函数调用(还在实验阶段!!)
>     clear ----------------------- 删除断点.
>     clearall -------------------- 删除多个断点.
>     condition (alias: cond) ----- 设置断点条件.
>     config ---------------------- 修改配置参数.
>     continue (alias: c) --------- 运行到断点或程序终止.
>     deferred -------------------- 在延迟调用的上下文中执行命令.
>     disassemble (alias: disass) - 反汇编程序.
>     down ------------------------ 将当前帧向下移动.
>     edit (alias: ed) ------------ 在$DELVE_EDITOR或$EDITOR中打开你所在的位置
>     exit (alias: quit | q) ------ 退出调试器.
>     frame ----------------------- 设置当前帧，或在不同的帧上执行命令.
>     funcs ----------------------- 打印函数列表.
>     goroutine ------------------- 显示或更改当前goroutine
>     goroutines ------------------ 列举程序goroutines.
>     help (alias: h) ------------- 打印帮助信息.
>     list (alias: ls | l) -------- 显示源代码.
>     locals ---------------------- 打印局部变量.
>     next (alias: n) ------------- 转到下一个源行.
>     on -------------------------- 在命中断点时执行命令.
>     print (alias: p) ------------ 计算一个表达式.
>     regs ------------------------ 打印CPU寄存器的内容.
>     restart (alias: r) ---------- 重启进程.
>     set ------------------------- 更改变量的值.
>     source ---------------------- 执行包含delve命令列表的文件
>     sources --------------------- 打印源文件列表.
>     stack (alias: bt) ----------- 打印堆栈跟踪信息.
>     step (alias: s) ------------- 单步执行程序.
>     step-instruction (alias: si)  单步执行一条cpu指令.
>     stepout --------------------- 跳出当前函数.
>     thread (alias: tr) ---------- 切换到指定的线程.
>     threads --------------------- 打印每个跟踪线程的信息.
>     trace (alias: t) ------------ 设置跟踪点.
>     types ----------------------- 打印类型列表
>     up -------------------------- 向上移动当前帧.
>     vars ------------------------ 打印包变量.
>     whatis ---------------------- 打印表达式的类型.
> 在命令前键入help来获得命令的完整文档，如help goroutine
> */
> 
> //常用操作总结
> 
> /*
> enter Delve会默认重复上一条命令
> 
> c continue 运行程序到断点或程序终止
> r restart 重启进程
> 
> b break 添加断点 
> 1.b main.main 在函数位置设置断点
> 2.b temp.go 12 在文件的行数设置断点
> 
> bp breakpoints 输出活动断点信息
> l list 显示从停止行起前后5行的源代码
> n next 转到下一个源行
> p print 计算一个表达式
> locals 打印局部变量
> s step 单步执行程序
> whatis 打印表达式类型
> 
> goroutine 显示或更改当前goroutine
> goroutine 				调用时不带参数，它将显示关于当前goroutine的信息。
> goroutine <id>			使用单个参数调用时，它将切换到指定的goroutine。
> goroutine <id> <command> 使用更多参数调用时，它将在指定的goroutine上执行命令。
> 
> goroutines 列举程序goroutines
> -u  显示用户代码中最顶层堆栈帧的位置
> -r  显示最顶层stackframe的位置(包括私有运行时函数中的帧)
> -g  显示创建goroutine的go指令的位置
> -s  显示起始函数的位置
> -t  显示goroutine的堆栈跟踪
> 
> args 打印函数参数
> stepout 跳出当前函数
> clear删除断点\clearall删除多个断点
> on 在命中断点时执行命令
> set 更改变量的值
> up向上移动当前帧/dowm向下移动当前帧
> exit (alias: quit | q)退出调试器
> */
> ```
>
> 