**mapReduce论文笔记**

---------------------------

**一、编程模型**

> 计算任务将一个键值对集合作为输入，并生成一个键值对集合作为输出。
>
> Map函数接收输入，生成中间键值对集合。
>
> MapReduce库将所有共用一个键的值组合，并将其传递给Reduce函数。
>
> Reduce接收一个中间键以及该键的值集合作为输入，将值进行合并。

```c
//map函数会返回一个单词加上它出现的次数的键值对
map(String key, String value):
// key: document name
// value: document contents
for each word w in value:
EmitIntermediate(w,"1");

//reduce函数会将该单词的出现次数统计在一起
reduce(String key, Iterator values):
// key: a word
// values: a list of counts
int result = 0;
for each v in values:
result += ParseInt(v);
Emit(AsString(result));

map(k1,v1)--> list(k2,v2)
reduce(k2,list(v2))--> list (v2)
```

**二、实现**

![](C:%5CUsers%5C18048%5CPictures%5C1-1608709719508.PNG)

> 对输入数据进行分片，可以实现map操作在多台机器上并行处理。
>
> 使用分区函数将map函数产生的中间key值划分为多个分区，这样Reduce操作也可以分布到多台机器上并行处理。
>
> 整个执行流程如下：
>
> 1. MapReduce库会先将输入文件切分为`M`个片段，通常每个片段的大小在16MB到64MB之间（具体大小可以由用户通过可选参数来进行指定）。接着，它会在集群中启动许多个程序副本。
> 2. 特殊程序master会给每个空闲的worker分配map任务或reduce任务。
> 3. 被分配了map任务的worker会读取相关的输入数据片段。它会从输入数据中解析出键值对，并将它们传入用户定义的`Map`函数中。`Map`函数所生成的中间键值对会被缓存在内存中。
> 4. 每隔一段时间，被缓存的键值对会被写入到本地硬盘，并通过分区函数分到R个区域内。这些被缓存的键值对在本地磁盘的位置会被传回`master`。`master`负责将这些位置转发给执行`reduce`操作的`worker`。
> 5. 当`master`将这些位置告诉了某个执行`reduce`的`worker`，该`worker`就会使用`RPC`的方式去从保存了这些缓存数据的`map worker`的本地磁盘中读取数据。当一个`reduce worker`读取完了所有的中间数据后，它就会根据中间键进行排序，这样使得具有相同键值的数据可以聚合在一起。之所以需要排序是因为通常许多不同的`key`会映射到同一个`reduce`任务中。如果中间数据的数量太过庞大而无法放在内存中，那就需要使用外部排序。
> 6. `reduce worker`会对排序后的中间数据进行遍历。然后，对于遇到的每个唯一的中间键，`reduce worker`会将该`key`和对应的中间`value`的集合传入用户所提供的`Reduce`函数中。`Reduce`函数生成的输出会被追加到这个`reduce`分区的输出文件中。
> 7. 当所有的`map`任务和`reduce`任务完成后，`master`会唤醒用户程序。此时，用户程序会结束对`MapReduce`的调用。
>
> master数据结构：
>
> 它保存了每个`Map`任务和每个`Reduce`任务的状态（闲置，正在运行，以及完成），以及非空闲任务的`worker`机器的`ID`。它将`map`任务所生成的中间文件区域的位置传播给`reduce`任务。对于每个完成的`map`任务，`master`会保存由`map`任务所生成的`R`个中间文件区域的位置和大小。当`map`任务完成后，会对该位置和数据大小信息进行更新。这些信息会被逐渐递增地推送给那些正在运行的`Reduce`工作。
>
> 容错：
>
> 1. worker故障
>
>    `master`会周期性`ping`下每个`worker`。如果在一定时间内无法收到来自某个`worker`的响应，那么`master`就会将该`worker`标记为`failed`。所有由该`worker`完成的`Map`任务都会被重设为初始的空闲（`idle`）状态。因此，之后这些任务就可以安排给其他的`worker`去完成。
>
>    当`worker`故障时，由于已经完成的`Map`任务的输出结果已经保存在该`worker`的硬盘中了，并且该`worker`已经无法访问，所以该输出也无法访问。因此，该任务必须重新执行。然而，已经完成的`Reduce`任务则无需再执行，因为它们的输出结果已经存储在全局文件系统中了。
>
>    当一个`Map`任务由`worker A`先执行，但因为`worker A`故障了，之后交由`worker B`来执行。所有执行`Reduce`任务的`woker`就会接受到这个重新执行的通知。任何还没有从`worker A`中读取数据的`Reduce`任务将从`worker B`中读取数据。
>
> 2. master故障
>
>    一个简单的解决好办法就是让`master`周期性的将上文所描述的数据结构写入磁盘，即`checkpoint`。如果这个`master`挂掉了，那么就可以从最新的`checkpoint`创建出一个新的备份，并启动`master`进程。
>
>    