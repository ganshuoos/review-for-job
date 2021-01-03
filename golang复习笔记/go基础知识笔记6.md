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

//反射
//Go语言中的变量是分为两部分的:
//类型信息：预先定义好的元信息。
//值信息：程序运行过程中可动态变化的。

//反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

//支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

//reflect包
//在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的
// 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

//TypeOf
//使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type）
//type name和type kind
//在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
//使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型
//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。\

//ValueOf
//reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func main() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}

//isNil()和isValid()
//isNil()
//func (v Value) IsNil() bool
//IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
//func (v Value) IsValid() bool
//IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
//IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

//结构体反射

```





