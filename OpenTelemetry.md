# OpenTelemetry

开源的可观测性框架和标准集合，用于对分布式系统进行监控，追踪和诊断


### 三类可观测指标

- `Traces`
  - 分布式追踪
    - 一次请求在多个服务中的调用链路和耗时
- `Metrics`
  - 指标，记录系统运行状态
    - `QPS`
    - `CPU`
- `Logs`
  - 记录离散事件并和 `Trace` 进行关联

##### 指标采集

使用 `Prometheus`

##### 日志采集

日志一般使用 `Sidecar` 代理的方式通过 `Agent` 将数据写入 `ElasticSearch`

### 架构

- 客户端（应用）
  - 挂载 `agent` 将采集信息上传 `Collector`
- `Collector`
  - 接收客户端上传数据，内部处理，输出到存储系统
- 数据存储



##### Collector

`Receiver` 用于接收客户端上报数据，除了 `agent` 还支持其他第三方的

`Exporter` 将 `Receiver` 收到的数据进行处理之后输出到不同组件

> `https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver` 查看所有支持的第三方 `Receiver`

`Receiver` 和 `Exporter` 可以自由组合实现任何逻辑并且不会影响到业务本身

