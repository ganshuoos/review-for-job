**golang切片**

> **切片**
>
> 切片三部分： 底层数组， 长度， 容量
>
> 切片存储在连续内存上（实际上是数组）
>
> **make 与 new**
>
> 使用make初始化会开辟一段内存，并初始化为对应的零值。
>
> <img src="golang%E5%88%87%E7%89%87.assets/make%E5%88%87%E7%89%87.PNG" style="zoom:50%;" />
>
> 使用new初始化字符串时，同样会分配这三个部分，但是new不负责底层数组的分配。new的返回值就是slice的起始地址。通过append的方式分配底层数组（不能直接访问， 底层数组为nil）。
>
> <img src="golang%E5%88%87%E7%89%87.assets/new%20%E5%88%9D%E5%A7%8B%E5%8C%96string.PNG" style="zoom:50%;" />
>
> **底层数组** （长度不可变， 多个切片可共用一个底层数组）
>
> <img src="golang%E5%88%87%E7%89%87.assets/%E5%BA%95%E5%B1%82%E6%95%B0%E7%BB%84-1610026910855.PNG" style="zoom:50%;" />
>
> **slice扩容规则**
>
> 1. 预估扩容后容量newCap
>
>    ​	
>
>    ```go
>    //oldCap 为旧容量， newCap为新容量, cap为需要扩容大小
>    if oldCap * 2 < cap {
>        newCap = cap
>    } else if oldLen < 1024 {
>        newCap = oldCap * 2
>    } else {
>        //循环
>        newCap = oldCap * 1.25
>    }
>    ```
>
>    
>
> 2. newCap个元素需要内存
>
>    ```go
>    //所需内存 = 预估容量 * 元素类型大小
>    //对齐原则！
>    //语言的内存管理模块会优先向操作系统申请常用规格大小的内存管理，申请内存时， 内存管理模块会匹配到足够大且最接近的规格
>    ```
>
>    
>
> 3. 匹配到合适的内存规格

