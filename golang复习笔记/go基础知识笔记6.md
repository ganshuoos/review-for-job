go基础知识笔记6

```go
//包
//一个包可以简单理解为一个存放.go文件的文件夹。 该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。
//一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
//包名可以不和文件夹的名字一样，包名不能包含 - 符号。
//包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

//包的可见性
//如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。

//包的导入
//要在代码中引用其他包的内容，需要使用import关键字导入使用的包。
import "包的路径"
//import导入语句通常放在文件开头包声明语句的下面。
//导入的包名需要使用双引号包裹起来。
//包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
//Go语言中禁止循环导入包。
import (
    "包1"
    "包2"
)

//自定义包名
//import 别名 "包的路径"
import "fmt"
import m "github.com/Q1mi/studygo/pkg_test"

func main() {
	fmt.Println(m.Add(100, 200))
	fmt.Println(m.Mode)
}

//匿名导入包
//如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。
import _ "包的路径"

//init()初始化函数
//在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是： init()函数没有参数也没有返回值。 init()函数在程序运行时自动被调用执行，不能在代码中主动调用它。

//init()函数执行顺序
//Go语言包会从main包开始检查其导入的所有包，每个包中又可能导入了其他的包。Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。在运行时，被最后导入的包会最先初始化并调用其init()函数.



//接口
//接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
//在Go语言中接口（interface）是一种类型，一种抽象的类型。interface是一组method的集合，是duck-type programming的一种体现。

//Go语言提倡面向接口编程。
//每个接口由数个方法组成，接口的定义格式如下：

type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
//接口名：使用type将接口定义为自定义的类型名。
//方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
//参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

//实现接口的条件
//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
// Sayer 接口
type Sayer interface {
	say()
}：
type dog struct {}
type cat struct {}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}
// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}

//值接收者和指针接收者实现接口的区别
//使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。
//此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。

//接口嵌套
// Sayer 接口
type Sayer interface {
	say()
}
// Mover 接口
type Mover interface {
	move()
}
// 接口嵌套
type animal interface {
	Sayer
	Mover
}

//空接口
//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
//空接口类型的变量可以存储任意类型的变量。

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
// 空接口作为map值
var studentInfo = make(map[string]interface{})
studentInfo["name"] = "沙河娜扎"
studentInfo["age"] = 18
studentInfo["married"] = false
fmt.Println(studentInfo)

//接口值
//一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

//类型断言
//x：表示类型为interface{}的变量
//T：表示断言x可能是的类型。
x.(T)
//该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}


```



