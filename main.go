package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pachirode/open-telemetry-demo/internal/provider"
	"github.com/pachirode/open-telemetry-demo/internal/serve"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
)

func run_grpc() {
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

	// var meter = otel.Meter("otel-demo")
	// counter, err := meter.Int64Counter("otel-demo.counter")
	if err != nil {
		slog.Error(err.Error())
	}
	tracer := otel.Tracer("otel-demo")

	// openfeature.AddHooks(otelhooks.NewTracesHook())
	serve.NewServer(tracer, nil)
	select {}
}

func run_http() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		return err
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())
	return err
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/rolldice/", rolldice)
	mux.HandleFunc("/rolldice/{player}", rolldice)

	handler := otelhttp.NewHandler(mux, "/")
	return handler
}

func main() {
	if err := run_http(); err != nil {
		log.Fatalln(err)
	}
}
