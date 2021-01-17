**go调试工具 – pprof**

> 参考：https://blog.jimmytinsley.com/2020/11/27/pprof%E8%B0%83%E8%AF%95%E6%8A%80%E5%B7%A7%E5%B0%8F%E7%BB%93/
>
> https://www.cnblogs.com/nickchen121/p/11517452.html
>
> pprof可以用来分析go程序(非Server)的运行时数据(`runtime/pprof`)和http server的运行时数据(`net/http/pprof`)。
>
> ```go
> //对于工具型应用
> import "runtime/pprof"
> //CPU性能分析
> pprof.StartCPUProfile(w io.Writer)
> pprof.StopCPUProfile()
> //内存性能分析
> pprof.WriteHeapProfile(w io.Writer)
> 
> //对于服务型应用
> //使用net/http
> import (
>     "net/http"
>     _ "net/http/pprof"   //加入这一行即可
> )
> 
> //对于gin框架而言
> import (
>     "github.com/gin-contrib/pprof"	//引入这个包
>     "github.com/gin-gonic/gin"	
> )
> 
> func main() {
>     route = gin.Default()
>     pprof.Register(route)	//注册一下路由
> }
> 
> //go tool pprof 指令
> go tool pprof [binary] [source]
> //binary 是应用的二进制文件，用来解析各种符号；
> //source 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址
> //top命令可以查看各项占用排名
> //list + 函数名 可以看到这个函数调用的代码上下文
> //go tool pprof --http :8081 http://your-host:port/debug/pprof/heap/ 可视化web界面例子
> 
> //比较两个时间点数据
> //1 dump内存
> curl -s http://your-host:port/debug/pprof/heap > previous.heap
> //2 dump内存
> curl -s http://your-host:port/debug/pprof/heap > current.heap
> //-base 参数进行比较
> go tool pprof --http :8081 --base previous.heap current.heap
> ```
>
> 