package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	v1 "github.com/blr-coder/book_grpc/api/v1"
	"github.com/blr-coder/book_grpc/internal/config"
	"github.com/blr-coder/book_grpc/internal/db"
	"github.com/blr-coder/book_grpc/internal/repositories"
	"github.com/blr-coder/book_grpc/internal/usecases"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func RunServer(ctx context.Context, config *config.Config, logger *logrus.Logger) error {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(
					logrus.NewEntry(logger),
				),
			),
		),
	)

	psqlDB, err := db.NewDB(
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Name,
		config.Postgres.User,
		config.Postgres.Password,
	)
	if err != nil {
		return fmt.Errorf("failed to init postgres: %v", err.Error())
	}

	bookRepository := repositories.NewBookRepository(psqlDB)
	bookUseCase := usecases.NewBookUseCase(bookRepository)
	bookGRPCServer := NewBookGRPCServer(bookUseCase)

	healthServer := health.NewServer()
	configureHealth(ctx, healthServer, psqlDB)

	v1.RegisterBookServer(grpcServer, bookGRPCServer)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	listener, err := net.Listen("tcp", config.BindAdr)
	if err != nil {
		return fmt.Errorf("failter to start listener: %v", err)
	}

	return grpcServer.Serve(listener)
}

func configureHealth(ctx context.Context, healthServer *health.Server, db *sqlx.DB) {
	healthServer.SetServingStatus("grpc.book.Book", grpc_health_v1.HealthCheckResponse_SERVING)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := db.PingContext(ctx); err != nil {
					healthServer.SetServingStatus("database", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
				} else {
					healthServer.SetServingStatus("database", grpc_health_v1.HealthCheckResponse_SERVING)
				}
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
}
