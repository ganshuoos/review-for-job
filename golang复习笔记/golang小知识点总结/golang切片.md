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
> <img src="D:%5Creview-for-job%5Cgolang%E5%A4%8D%E4%B9%A0%E7%AC%94%E8%AE%B0%5Cgolang%E5%B0%8F%E7%9F%A5%E8%AF%86%E7%82%B9%E6%80%BB%E7%BB%93%5Cimage%5Cmake%E5%88%87%E7%89%87.PNG" style="zoom:50%;" />
>
> 使用new初始化字符串时，同样会分配这三个部分，但是new不负责底层数组的分配。new的返回值就是slice的起始地址。通过append的方式分配底层数组（不能直接访问， 底层数组为nil）。
>
> <img src="D:%5Creview-for-job%5Cgolang%E5%A4%8D%E4%B9%A0%E7%AC%94%E8%AE%B0%5Cgolang%E5%B0%8F%E7%9F%A5%E8%AF%86%E7%82%B9%E6%80%BB%E7%BB%93%5Cimage%5Cnew%20%E5%88%9D%E5%A7%8B%E5%8C%96string.PNG" style="zoom:50%;" />
>
> **底层数组** （长度不可变， 多个切片可共用一个底层数组）
>
> <img src="golang%E5%88%87%E7%89%87.assets/%E5%BA%95%E5%B1%82%E6%95%B0%E7%BB%84-1610026910855.PNG" style="zoom:50%;" />

