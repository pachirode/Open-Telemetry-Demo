package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/pachirode/open-telemetry-demo/internal/provider"
	"github.com/pachirode/open-telemetry-demo/internal/serve"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := provider.InitTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.Error(fmt.Sprintf("Error shutting down provider provider: %v", err))
		}
	}()

	//mp := provider.InitMeterProvider()
	//defer func() {
	//	if err := mp.Shutdown(context.Background()); err != nil {
	//		slog.Error(fmt.Sprintf("Error shutting down meter provider: %v", err))
	//	}
	//}()

	err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))
	if err != nil {
		slog.Error(err.Error())
	}

	//var meter = otel.Meter("otel-demo")
	//counter, err := meter.Int64Counter("otel-demo.counter")
	if err != nil {
		slog.Error(err.Error())
	}
	tracer := otel.Tracer("otel-demo")

	//openfeature.AddHooks(otelhooks.NewTracesHook())
	serve.NewServer(tracer, nil)
	select {}
}
