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
>
>    依赖于`map`任务和 `reduce`任务输出的原子性提交，即每个正在执行的任务会将它的输出写入到私有的临时文件中去。每个`Reduce`任务会生成这样一个文件，每个`Map`任务则会生成`R`个这样的文件（一个`Reduce`任务对应一个文件）。当一个`map`任务完成时，该`map`任务对应的`worker`会向`master`发送信息，该信息中包含了`R`个临时文件的名字。如果`master`从一个已经完成的`map`工作的`worker`处又收到这个完成信息，`master`就会将该信息忽略。否则，它会将这`R`个文件名记录在`master`的数据结构中。
>
>    当`Reduce`任务完成时，`reduce worker`会以原子的方式将临时输出文件重命名为最终输出文件。如果多台机器执行同一个`reduce`任务，那么对同一个输出文件会进行多次重命名。我们依赖于底层文件系统所提供的原子性重命名操作来保证最终的文件系统状态仅包含一个`Reduce`任务所产生的数据。
>
> 地区性：
>
> `GFS`将每个文件分割为许多`64MB`大小的区块（`Block`），并且会对每个区块保存多个副本（分散在不同的机器上，通常是`3`个副本）。`MapReduce`的`master`在调度`Map`任务时会考虑`输入数据文件`的位置信息。尽量在包含该相关输入数据的拷贝的机器上执行`Map`任务。如果任务失败，`master`会尝试在保存输入数据副本的邻近机器上执行`Map`任务（例如，将任务交由与包含数据的机器在同一网络交换机下的`woker`机器去执行）。当一个集群中大部分`worker`机器都在执行`MapReduce`操作时，大部分输入数据会在本地进行读取，这样就不会消耗网络带宽。
>
>  任务粒度：
>
> `master`必须执行`O(M+R)`次调度，并且在内存中保存`O(M*R)`个状态（这对于内存的使用率来说影响还是比较小的，在`O(M*R)`个状态中，每个状态保存的每对`Map`任务/`Reduce`任务的数据大小为`1 byte`）。	在`MapReduce`计算中，我们常用的M大小为`200000`，R为`5000`，用到的`worker`机器数量为`2000`。
>
> 备用任务（落伍者）
>
> 当一个`MapReduce`计算接近完成时，`master`会调度一个备用（`backup`）任务来执行剩下的处于正在执行中（`in-progress`）的任务。无论是这个主任务还是这个备用任务完成了，我们都会将这个任务标记为完成。我们对这个机制进行了调优。通常情况下，它只会比正常操作多占几个百分点的计算资源。我们发现这样做能够显著减少运行大型`MapReduce`计算时所要花费的时间。



**四、改进**

> 分区函数
>
> 默认情况下，分区函数使用的是哈希进行取模（例如：`hash(key) mod R`）进行分区。这样能够生成非常均匀的数据分区。但是在某些情况下，通过向其他的一些分区函数传入`key`来进行分区会非常有用。比如，有时候输出的`key`为`URL`，我们希望每个主机的所有条目存放在同一个输出文件中。为了支持类似的情况，使用`MapReduce`库的用户可以提供一个特殊函数。例如，使用 **“hash(Hostname(key)) mod R”**作为分区函数，就可以把所有来自同一个主机的`URL`保存在同一个输出文件中。
>
> 顺序保证
>
> 在给定的分区中，我们保证中间键值对的处理顺序是根据`key`的大小进行升序排列。
>
> Combiner函数
>
> 我们允许用户提供一个可选的`Combiner`函数，在数据通过网络发送之前，可以通过该函数将数据进行部分合并。
>
> `Combiner`函数会在每台执行`Map`任务的机器上执行一次。通常情况下，`Combiner`函数和`Reduce`函数的实现代码是一样的。`Reduce`函数和`Combiner`函数唯一的区别在于`MapReduce`库处理函数输出上面会有所不同。`Reduce`函数的输出会被写入一个最终输出文件中，而`Combiner`函数的输出会被写入一个要发送给`reduce`任务的中间文件中。
>
> 输入输出数据
>
> 跳过损坏的记录
>
> 在处理大型数据集的时候。我们提供了一种可选的执行模式，当`MapReduce`库检测到某些记录绝对会引起崩溃的时候，它就会跳过这些记录，以此来保证计算进度的推进。
>
> 每个`worker`进程都会通过一个`handler`来捕获内存段异常（segmentation violation）和总线错误（bus error）。在调用用户的`Map`或`Reduce`操作前，`MapReduce`库会用一个全局变量来保存参数序号。如果用户代码产生了一个`signal`，`singnal handler`就会在发送最后一个`UDP`包时，在该`UDP`包中放入该序号并发送给`MapReduce master`的序号。当`master`检测到在某条记录上有多次故障的时候，当它发出相应的`Map`或者`Reduce`任务重新执行时，它就表示应该跳过这条记录。、
>
> 状态信息
>
> 在`master`中，有一个内置的`HTTP`服务器，它可以用来展示一组状态信息页面。状态页面会显示计算进度，例如：已经完成的任务数量、正在执行的任务数量、输入的字节数、中间数据的字节数、输出的字节数、处理率等等。
>
> 计数器
>
> `MapReduce`库提供了计数器机制，它能用来统计不同活动的发生次数。为了使用这种机制，用户需要创建一个名为`Counter`的对象，然后在`Map`和`Reduce`函数中以正确的方式增加`counter`的数字。
>
> ```c
> Count* uppercase;
> uppercase = GetCounter("uppercase");
> 
> map(String name, String contents) :
> for each word w in contents:
> if(isCapitalized(w)):
> uppercase->Increment();
> EmitIntermediate(w,"1");
> ```
>
> 每隔一段时间，这些`counter`值会从每个`worker`机器传递给`master`（附加在`ping`的应答包中进行传递）。当`MapReduce`操作完成时，`master`会将这些已经成功完成的`map`和`reduce`任务中返回的`counter`的值聚合在一起，并将它们返回给用户代码。当前的`counter`值会显示在`master`的状态页面中。这样，人们就可以看到当前计算的进度。当聚合这些`counter`的值时，`master`会去掉那些重复执行的相同`map`或者`reduce`操作的次数，以此避免重复计数。

