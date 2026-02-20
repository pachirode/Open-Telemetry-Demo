参考链接 https://opentelemetry.io/zh/docs/languages/python

# Python

### 安装

`opentelemetry-distro` 库，包含 `API`, `SDK`, `opentelemetry-bootstrap` 和 `opentelemetry-instrumnet`

```bash
pip install opentelemetry-distro
```

### 启动

```bash
export OTEL_PYTHON_LOGGING_AUTO_INSTRUMENTATION_ENABLED=true
opentelemetry-instrument \
    --traces_exporter console \
    --metrics_exporter console \
    --logs_exporter console \
    --service_name dice-server \
    flask run -p 8080
```

> 自动插桩依赖于 `opentelemetry-instrumentation-flask`，必须使用 `opentelemetry-instrument` 启动

### 结果导出

```bash
pip install opentelemetry-exporter-otlp

# 设置 OTEL 服务
$env:OTEL_EXPORTER_OTLP_ENDPOINT="http://192.168.52.1:4317"
$env:OTEL_TRACES_EXPORTER="otlp"
$env:OTEL_METRICS_EXPORTER="otlp"
$env:OTEL_LOGS_EXPORTER="otlp"
```

