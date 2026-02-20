# Trace 分布式追踪

一次完整请求在分布式系统的端到端执行路径

# Tracer Provider

`Trace` 工厂，大多数应用中只初始化一次，其生命周期和应用的生命周期一致
初始化的第一步为 `Resource` 和 `Explorter` 的初始化，有些 `SDK` 中以及初始化全局的 `Tracer Provider`

### Exporter

将链路信息发送给消费者，消费者可以是标准输出，`OpenTelemetry Collector` 或者其他后端

### Trace

一个 `Trace` 通常对应一次请求，其本身不直接记录动作，而是由多个 `Span` 组成
每个 `Trace` 的有一个全局唯一的 `TraceID`，属于该请求的 `Span` 都需要携带这个 `TraceID`

### Span

通常情况下调用本地函数不会再添加 `Span`，如果本地函数特别重要可以使用注解

##### 包含信息

- 基本信息
    - `traceId`
    - `spanName`
    - `spanId`
    - `pranetId`
    - 定义调用的主次，如果是并行调用则这个字段相同
    - 开始时间
    - 结束时间
- 内置信息
  - `status`
    - `Error`
    - `Ok`
- 类型
  - `Client`
  - `Server`
  - `Internal`

##### 上下文

`Span` 上下文是不可变对象是 `Span` 的一部分，包含一些链路的信息，可以用来创建 `Span` 链接


### Baggage

`Span` 中的上下文，可以将我们想要的数据存放在其中，这样任意一个 `Span` 都可以读到
自动插桩会在大多数网络中包含 `Baggage` 可能会导致信息泄漏
