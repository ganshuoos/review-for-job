golang 基础知识笔记（二）

----------------------------------

一、数组

```go
//在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
//数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。
var 数组变量名 [元素数量]T
var arr [3]int

var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型

//数组初始化
//方法一：初始化数组时可以使用初始化列表来设置数组元素的值
var cityArray = [3]string{"北京", "上海", "深圳"}
//方法二：让编译器根据初始值的个数自行推断数组的长度
var cityArray = [...]string{"北京", "上海", "深圳"}
//方法三：使用指定索引值的方式来初始化数组
a := [...]int{1: 1, 3: 5}
fmt.Println(a)                  // [0 1 0 5]

//二维数组
a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
//多维数组只有第一层可以使用...来让编译器推导数组长度。
//支持的写法
a := [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
//数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
//[n]*T表示指针数组，*[n]T表示数组指针 。
```

二、切片

```go
//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。
//切片是一个引用类型，它的内部结构包含地址、长度和容量。
//切片的底层就是一个数组

//切片声明
//name:表示变量名
//T:表示切片中的元素类型
var name []T

//求切片的长度和容量
len(slice): 求长度 cap(slice): 求容量

//对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

s:[2 3] len(s):2 cap(s):4
s2:[5] len(s2):1 cap(s2):1

//完整切片表达式
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}
t:[2 3] len(t):2 cap(t):4

//使用make()函数构造切片
make([]T, size, cap)
//T:切片的元素类型
//size:切片中元素的数量
//cap:切片的容量 可省略

//切片本质
//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

//判断切片是否为空
//使用len(slice) == 0判断，slice == nil 表示为nil切片
//nil切片的特点 len 和 cap都为0

//切片不能直接比较
//切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。

//切片的赋值拷贝
//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容

//append()方法 为切片添加元素
s2 := []int{5, 6, 7}  
	s = append(s, s2...)    // [1 2 3 4 5 6 7]
}
//通过var声明的零值切片可以在append()函数直接使用，无需初始化。
var s []int
s = append(s, 1, 2, 3)

//切片扩充策略
//每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
//append()函数将元素追加到切片的最后并返回该切片。
//切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
//append()函数还支持一次性追加多个元素。
citySlice = append(citySlice, "上海", "广州", "深圳")

//扩充策略源码
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
//先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
//否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
//否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
//如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

//使用copy()函数复制切片
//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
copy(destSlice, srcSlice []T)
//srcSlice: 数据来源切片
//destSlice: 目标切片

//从切片中删除元素
a = append(a[:2], a[3:]...)
//要从切片a中删除索引为index的元素，操作方法是
a = append(a[:index], a[index+1:]...)

//切片与数组的区别
```

