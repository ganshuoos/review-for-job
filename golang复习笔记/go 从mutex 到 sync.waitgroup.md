#### go 从 mutex 到 sync.WaitGroup

参考：	https://blog.csdn.net/qq_43313035/article/details/98659418

​				https://www.cnblogs.com/buyicoding/p/12082162.html

​				https://zhuanlan.zhihu.com/p/75263302

​				https://www.peihuan.net/2019/03/05/Treap%E6%A0%91/

​				https://blog.csdn.net/qq_37005831/article/details/111112319?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.control

> > ##### 锁机制
> >
> > ##### 悲观锁
> >
> > - 假定会发生并发冲突，屏蔽一切可能违反数据完整性的操作。
> > - 悲观锁的实现，往往依靠底层提供的锁机制。
> > - 悲观锁会导致其它所有需要锁的线程挂起，等待持有锁的线程释放锁。
> >
> > ##### 乐观锁
> >
> > - 假设不会发生并发冲突，每次不加锁而是假设没有冲突而去完成某项操作，只在提交操作时检查是否违反数据完整性。
> > - 如果因为冲突失败就重试，直到成功为止。
> > - 乐观锁大多是基于数据版本记录机制实现。
> > - 为数据增加一个版本标识
> > - 乐观锁的缺点是不能解决脏读的问题
> >
> > ##### 乐观锁与悲观锁总结
> >
> > - 并发量不大且不允许脏读 – > 悲观锁
> > - 并发量非常大 – > 乐观锁
> >
> > ##### 锁的公平性
> >
> > - 在锁获取的过程中，先进行锁获取的线程是否比后续的线程更先获得锁，如果是则就是公平锁：多个线程按照获取锁的顺序依次获得锁，否则就是非公平性
> >
> > 
>
> > ##### CAS （Compare & Set/Compare & Swap）
> >
> > - CAS操作包含三个操作数——内存位置（V）、预期原值（A）、新值(B）
> > - CAS有效地说明了“我认为位置V应该包含值A；如果包含该值，则将B放到这个位置；否则，不要更改该位置，只告诉我这个位置现在的值即可。
> >
> > ##### CAS操作
> >
> > - **比较 A 与 V 是否相等**
> > - **如果比较相等，将 B 写入 V**
> > - **返回操作是否成功**
>
> > ##### Mutex
> >
> > ```go
> > //mutexLocked对应右边低位第一个bit。值为1，表示锁被占用。值为0，表示锁未被占用。
> > //mutexWoken对应右边低位第二个bit。值为1，表示打上唤醒标记。值为0，表示没有唤醒标记。
> > //mutexStarving对应右边低位第三个bit。值为1，表示锁处于饥饿模式。值为0，表示锁存于正常模式。
> > //mutexWaiterShift是偏移量。它值为3。用法是state>>=mutexWaiterShift之后，state的值就表示当前阻塞等待锁的goroutine个数。最多可以阻塞2^29个goroutine。
> > 
> > 
> > type Mutex struct {
> > 	// [阻塞的goroutine个数, starving标识, woken标识, locked标识]
> > 	state int32		//通过state来进行锁的计数
> > 	sema  uint32	//通过sema来实现排队
> > }
> > 
> > const (
> > 	mutexLocked = 1 << iota // mutex is locked
> > 	mutexWoken    
> > 	mutexStarving 
> > 	mutexWaiterShift = iota
> > )
> > ```
> >
> > ##### 锁模式
> >
> > |          | 描述                                                         | 公平性 |
> > | -------- | ------------------------------------------------------------ | ------ |
> > | 正常模式 | 正常模式下所有的goroutine按照FIFO的顺序进行锁获取,被唤醒的goroutine和新请求锁的goroutine同时进行锁获取，通常新请求锁的goroutine更容易获取锁 | 否     |
> > | 饥饿模式 | 饥饿模式所有尝试获取锁的goroutine进行等待排队，新请求锁的goroutine不会进行锁获取，而是加入队列尾部等待获取锁 | 是     |
> >
> > ##### 自旋
> >
> > 自旋就是CPU空转一定的时钟周期
> >
> > 自旋需要满足以下条件：
> >
> > 1. 锁已被占用，并且锁不处于饥饿模式。
> > 2. 积累的自旋次数小于最大自旋次数（active_spin=4）。
> > 3. cpu核数大于1。
> > 4. 有空闲的P。
> > 5. 当前goroutine所挂载的P下，本地待运行队列为空。
> >
> > ```go
> > const 	active_spin     = 4
> > func sync_runtime_canSpin(i int) bool {
> > 	// 自旋次数不能大于 active_spin(4) 次
> > 	// cpu核数只有一个，不能自旋
> > 	// 没有空闲的p了，不能自旋
> > 	if i >= active_spin || ncpu <= 1 || gomaxprocs <= int32(sched.npidle+sched.nmspinning)+1 {
> > 		return false
> > 	}
> > 	// 当前g绑定的p里面本地待运行队列不为空，不能自旋
> > 	if p := getg().m.p.ptr(); !runqempty(p) {
> > 		return false
> > 	}
> > 	return true
> > }
> > ```
> >
> > 
> >
> > ##### 唤醒
> >
> > 唤醒标志主要用于标识当前尝试获取goroutine是否有正在处于唤醒状态的。
> >
> > 若新goroutine申请锁，锁被占用且自身满足自旋条件时，自旋并设置Woken标记。当占用锁的goroutine释放锁时，查看woken标记。此时若有标记，即使阻塞队列不为空，也不唤醒。直接return， 使自旋goroutine更可能抢到锁！
> >
> > 释放锁时，检查Woken标记为空。而阻塞队列里有goroutine需要被唤醒。那么在唤醒时，同时标记锁Woken。原来没有Woken标记，为什么在唤醒一个goroutine要主动标记呢？目的是保证锁公平。设置Woken标记后，state就肯定不为零。此时新来的竞争者，在执行Lock()的fast-path时会失败，接下来就只能乖乖排队了。
> >
> > ```go
> > // 唤醒一个阻塞的goroutine，并把锁的Woken标记设置上
> > new = (old - 1<<mutexWaiterShift) | mutexWoken
> > 
> > func (m *Mutex) Lock() {
> > 	// Fast path: grab unlocked mutex.
> > 	// Woken标记设置后，这里的CAS就会为false
> > 	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
> > 		// ...
> > 		return
> > 	}
> >   // 接下来在阻塞里排队
> > }
> > ```
> >
> > ##### 锁尽量公平
> >
> > 尽量公平的结果就是阻塞的竞争者被唤醒后，也要与(正在自旋的)新竞争者抢夺锁资源。
> >
> > 1. 在锁释放时，主动设置Woken标记，防止新的竞争者轻易抢到锁。
> > 2. 竞争者进阻塞队列策略不一样。新的竞争者，抢不到锁，就排在队列尾部。先来竞争者，从队列中被唤醒后，还是抢不到锁，就放在队列头部。
> > 3. 任何竞争者，被阻塞等待的时间超过指定阀值(1ms)。锁就转为饥饿模式。这时锁释放时会唤醒它们，手递手式把锁资源给它们。别的竞争者（包括新来的）都抢不到。直到把饥饿问题解决掉。
> >
> > ##### 阻塞和唤醒机制semacquire和semrelease
> >
> > go的runtime有一个全局变量semtable，它放置了所有的信号量。
> >
> > 如果addr大于1，并且通过CAS减一成功，那就说明获取信号量成功。不用阻塞。
> >
> > 否则，semacquire1会在semtable数组中找一个元素和它对应上。每个元素都有一个root，这个root是Treap树。
> >
> > 最后addr变成一个树节点，这个树节点，有自己的一个队列，专门放被阻塞的goroutine。叫它阻塞队列吧。
> >
> > 这个阻塞队列是个双端队列，头尾都可以进。
> >
> > semacquire1把当前goroutine相关元数据放进阻塞队列之后，就挂起了。
> >
> > semrelease1是给addr CAS加一。
> >
> > 如果坚持发现当前addr上有阻塞的goroutine时，就取一个出来，唤醒它，让它自己再去semacquire1。这是handoff为false的情况。
> >
> > 但handoff为true的话，就尝试手递手地把信号量送给这个goroutine。等于说goroutine不用再自己去抢了，因为自己再去抢有可能抢不到。
> >
> > 最后semrelease1会把取出来的这个goroutine挂在当前P的本地待运行队列尾部，等待调度执行。
>
> > ##### sync.WaitGroup
> >
> > ```go
> > type WaitGroup struct {
> > 	noCopy noCopy
> > 	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
> > 	// 64-bit atomic operations require 64-bit alignment, but 32-bit
> > 	// compilers do not ensure it. So we allocate 12 bytes and then use
> > 	// the aligned 8 bytes in them as state, and the other 4 as storage
> > 	// for the sema.
> > 	state1 [3]uint32
> > }
> > 
> > //state1是个长度为3的数组，其中包含了state和一个信号量，而state实际上是两个计数器.
> > //counter： 当前还未执行结束的goroutine计数器
> > //waiter count: 等待goroutine-group结束的goroutine数量，即有多少个等候者
> > //semaphore: 信号量
> > 
> > //Add(delta int): 将delta值加到counter中
> > //Wait()： waiter递增1，并阻塞等待信号量semaphore
> > //Done()： counter递减1，按照waiter数值释放相应次数信号量
> > 
> > 
> > ```
> >
> > 

