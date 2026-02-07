package serve

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pachirode/open-telemetry-demo/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	api.UnimplementedGreeterServer
	tracer  trace.Tracer
	counter metric.Int64Counter
}

func NewServer(tracer trace.Tracer, counter metric.Int64Counter) {
	s := &server{
		tracer:  tracer,
		counter: counter,
	}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	api.RegisterGreeterServer(serve, s)

	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Printf("serving on port 8080")
	}

	quite := make(chan os.Signal)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quite
		log.Printf("Received termination signal, shutting down")
		os.Exit(0)
	}()

	log.Printf("Serving on port 8080")
}

func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	//defer s.counter.Add(ctx, 1)
	metaData, _ := metadata.FromIncomingContext(ctx)
	name, _ := os.Hostname()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("request.name", in.Name))
	s.span(ctx)
	return &api.HelloReply{Message: fmt.Sprintf("hostname:%s, in:%s, md:%v", name, in.Name, metaData)}, nil
}

func (s *server) span(ctx context.Context) {
	ctx, span := s.tracer.Start(ctx, "span")
	defer span.End()
	slog.Info("Create Span")
}
