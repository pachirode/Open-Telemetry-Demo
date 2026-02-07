# Trace 分布式追踪

一次完整请求在分布式系统的端到端执行路径

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

