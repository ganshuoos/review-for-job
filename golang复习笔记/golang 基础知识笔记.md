golang 基础知识笔记

-------------------------------------------------------------

一、变量与常量

```go
//标准声明
var 变量名 变量类型
var name string

//批量声明
var (
	a string
    b int
    c bool
)

//变量初始化
var 变量名 类型 = 表达式
var age int = 18

//短变量声明
m := 10

//匿名变量
//匿名变量不占用命名空间，不会分配内存，不存在重复声明

//常量
const pi = 3.14
const (
	pi = 3.14
    e = 2.71
)
const (
	n1 = 100
    n2 //100
    n3 //100
)

//itoa 常量计数器
//itoa 在const关键字出现时将被重置为0 const中每新增一行常量声明将使iota计数一次
const (
	n1 = iota //0
	n2        //1
	_
	n4        //3
)

const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	n4        //3
)
const n5 = iota //0

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)
```

二、基本数据类型

```go
//整型
//无符号整型： uint8(byte) uint16(short) uint32 uint64(long)
//有符号整型：int8 int16 int32 int64
//特殊整型： int uint 由32bit/64bit系统决定 uintptr 无符号整形指针

//注 常用整型常量：
//math.MaxInt32
//math.MinInt32

//数字字面量语法
//0b => 二进制
//0o => 八进制 
//0x => 十六进制

//浮点型
//支持两种浮点型数：float32和float64

//复数
//复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
//var c1 complex64
//c1 = 1 + 2i
//var c2 complex128
//c2 = 2 + 3i

//布尔值
//true false
//Go 语言中不允许将整型强制转换为布尔型.

//字符串
//注：字符串不可变！
//多行字符串：
// s1 := `1
//		  2`

//字符串常用操作
len(str)	//长度
str1 + str2 || fmt.Sprintf("%s%s", str1, str2)	//拼接
strings.Split	//分割
strings.contains	//包含
strings.HasPrefix("str")	//前缀
strings.HasSuffix("str")	//后缀
strings.Index('s')	//自左索引
strings.LastIndex('s') //自右索引
strings.Join(a[]string, sep string)	//join
strings.Join([]string{"1", "2"}, "$$")

//byte 与 rune
//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符
//rune类型，代表一个 UTF-8字符 rune类型实际是一个int32

//由于字符串不可变 修改字符串需先将其转化为[]rune或[]byte，完成后再转换为string
s1 := "big"
// 强制类型转换
byteS1 := []byte(s1)
byteS1[0] = 'p'

s2 := "白萝卜"
runeS2 := []rune(s2)
runeS2[0] = '红'

//强制类型转换（没有隐式类型转换）
T(表达式)
```

三、流程控制

```go
//if - else
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}

//for (没有while)
for 初始语句;条件表达式;结束语句{
    循环体语句
}

//for range 遍历数组、切片、字符串、map 及通道（channel）
//数组、切片、字符串返回索引和值。
//map返回键和值。
//通道（channel）只返回通道内的值。

//switch case
//1.每个switch语句只能有一个default
//2.一般满足条件后不继续执行，fallthrough语法可以执行满足条件的case的下一个case
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

//goto无条件跳转
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
    
//break 
//break语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上
func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
    
//continue
//continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
//在 continue语句后添加标签时，表示开始标签对应的循环。
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
```

