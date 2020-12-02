### 各语言异常处理的演进

- C	

  单返回值，一般通过传递指针作为入参，返回值为int表示成功或失败

- C++

  引入了exception，但是无法知道被调用方会抛出什么异常

- Java

  引入了checked exception，方法的所有者必须声明，调用者必须处理。异常的严重性由调用者来区分

- Go

  使用多个返回值和一个简单的约定，在返回值中带上实现了error接口的对象，由调用者处理

  Go panic 意味着程序运行失败，会退出

### Error Type

- Sentinel Error

  **预定义的特定错误**

  调用方必须使用 == 将结果与预先声明的值进行比较，如果返回方需要返回更多上下文信息，必然会破坏调用方的判等操作

  在两个包之间建立了源代码依赖关系

  sentinel 值是最不灵活的错误处理策略，应该尽可能避免sentinel errors

- Error Types

  **实现了Error接口的自定义类型**

  可以包装底层的错误以提供更多上下文

  调用者要使用类型断言和类型switch，就要让自定义的error 变为public，这会导致和调用者产生强耦合

  应该避免使用，或者至少避免它们作为公共API的一部分

- Opaque Errors

  **不透明错误处理**

  只需返回错误而不假设其内容

  作为调用者，关于操作的结果所知道的就是成功或者失败，没有能力看到错误的内部

  可以断言错误实现了特定的行为，而不是断言错误是特定的类型或值

  它要求代码和调用者之间的耦合最少，是最灵活的错误处理策略

### Handling Error

#### Wrap errors

通过使用pkg/errors包，可以向错误值添加上下文，这种方式既可以由人也可以由机器检查

- 如果和其他库进行协作，使用errors.Wrap或者errors.Wrapf保存堆栈信息。同样适用于和标准库协作的时候
- 直接返回错误，而不是每个错误产生的地方到处打日志。 
- 在程序的顶部或者是工作的goroutine 顶部(请求入口)，使用%+v把堆栈详情记录

#### Unwrap

包含另一个错误的错误可以实现Unwrap方法来返回所包含的底层错误，主要用于自定义错误包含底层错误时

如果 e1.Unwrap() 返回e2，那么我们说e1 包装e2，可以展开e1 以获得e2

#### Is & As

go1.13 errors 包中包含两个用于检查错误的新函数：Is 和 As

errors.Is函数的行为类似于(sentinel error)的判等操作，而errors.As函数的行为类似于类型断言(type assertion)

在处理包装错误(包含其他错误的错误)时，这些函数会考虑错误链中的所有错误





### 问题文档

第三课：	https://shimo.im/docs/vr6yDVPxRxXGKRDd

第四课：	https://shimo.im/docs/R6gP8qyvWqJrgRCk

相关文章：https://shimo.im/docs/GYvDrQT8qW8RgkGY

全部链接：https://shimo.im/docs/MqLqgJAFXigaXvPt/read