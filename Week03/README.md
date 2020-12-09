## go并发编程

### goroutine

#### G-P-M

相关文章：https://time.geekbang.org/column/article/39841

G(goroutine)代表go语言的用户级线程--协程，M(machine)代表系统级线程，P(processor)用于对接G和M

#### 使用goroutine注意点

- 把并发交给调用者
- 搞清楚什么时候退出
- 能够控制goroutine正常退出

### Memory model

#### MemoryReordering

相关文章：https://blog.csdn.net/qcrao/article/details/92759907

**软件或硬件系统可以根据其对代码的分析结果，一定程度上打乱代码的执行顺序，以达到其不可告人的目的。软件指的是编译器，硬件指的是CPU，不可告人的目的就是：减少程序指令数、最大化提高CPU利用率**

- CPU重排：现代 CPU 为了“抚平” 内核、内存、硬盘之间的速度差异，搞出了各种策略，例如三级缓存等

- 编译器重排：编译器会根据代码的上下文语义进行指令重排，可能会改变代码的执行顺序

#### Happens Before

相关文章：https://docs.studygolang.com/ref/mem

**在读操作和写操作之间不允许有其它的写操作发生**

### pachage sync

#### sync

- Mutex

  使用互斥锁的注意事项如下：

  - 不要重复锁定互斥锁
  - 不要忘记解锁互斥锁，必要时使用defer语句
  - 不要对尚未锁定或者已解锁的互斥锁解锁
  - 不要在多个函数之间直接传递互斥锁

- RWMutex

  - 同时只能有一个 goroutine 能够获得写锁定
  - 同时可以有任意多个 gorouinte 获得读锁定
  - 同时只能存在写锁定或读锁定（读和写互斥）

- WaitGroup

  用于等待一组 goroutine 结束。使用时wg.Add(num)和wg.Wait()方法要写在所管控的goroutine外部

- Cond

  - 实现一个条件变量，即等待或宣布事件发生的 goroutines 的会合点，用于阻塞和唤醒协程。
  - 在调用 Signal 或者 Broadcast 之前，要确保目标协程处于 Wait 阻塞状态，不然会出现死锁问题。
  - 和 Java 的等待唤醒机制很像，它的三个方法 Wait、Signal、Broadcast 就分别对应 Java 中的 wait、notify、notifyAll。

- Pool

  可以作为临时对象的保存和复用的集合。适用与无状态的对象的复用，而不适用与如连接池之类的

- Once

  可以使得函数多次调用只执行一次。适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景

- atomic

  原子操作在进行的过程中是不允许中断的，这会由 CPU 提供芯片级别的支持。

  atomic.Value可以用于实现Copy-On-Write，写时复制一份数据用于读

- errgroup

  - errgroup 可以捕获和记录子协程的错误（只能记录最先出错的协程的错误）
  - errgroup 可以控制协程并发顺序。确保子协程执行完成后再执行主协程
  - errgroup 可以使用 context 实现协程撤销或者超时撤销。子协程中使用 ctx.Done()来获取撤销信号

### context

可以跟踪协程，只有跟踪到每个协程，才能更好地控制它们。context的显式传递使得跨 API 边界的请求范围元数据、取消信号和截止日期很容易传递给处理请求所涉及的所有goroutine

要更好地使用 Context，有一些使用原则需要尽可能地遵守。

- Context 不要放在结构体中，要以参数的方式传递。
- Context 作为函数的参数时，要放在第一位，也就是第一个参数。
- 要使用 context.Background() 函数生成根节点的 Context，也就是最顶层的 Context。
- Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。
- Context 多协程安全，可以在多个协程中放心使用。

### channel

#### 无缓冲通道

容量是 0，不能存储任何数据。 只起到传输数据的作用，数据并不会在 channel 中做任何停留。

必须同时有接收和发送的goroutine，否则会阻塞，**本质是为了保证同步**。

#### 缓冲通道

接收和发送是非同步的

- 有缓冲 channel 的内部有一个缓冲队列
- 发送操作是向队列的尾部插入元素，如果队列已满则阻塞等待，直到另一个 goroutine 执行接收操作释放队列的空间
- 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行发送操作插入新的元素
- 如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值