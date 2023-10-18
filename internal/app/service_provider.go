package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/romanfomindev/microservices-auth/internal/client/db"
	"github.com/romanfomindev/microservices-auth/internal/client/db/pg"
	"github.com/romanfomindev/microservices-auth/internal/client/db/transaction"
	"github.com/romanfomindev/microservices-auth/internal/closer"
	"github.com/romanfomindev/microservices-auth/internal/config"
	"github.com/romanfomindev/microservices-auth/internal/config/env"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user"
	"github.com/romanfomindev/microservices-auth/internal/services"
	userService "github.com/romanfomindev/microservices-auth/internal/services/user"
)

type serviceProvider struct {
	pgPool         *pgxpool.Pool
	dbClient       db.Client
	pgConfig       config.PGConfig
	txManager      db.TxManager
	grpcConfig     config.GRPCConfig
	userRepository repositories.UserRepository
	userService    services.UserService
	userHandlers   *handlers.UserV1Handlers
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.pgConfig.DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repositories.UserRepository {
	if s.userRepository == nil {
		repo := user.NewUserRepository(s.DBClient(ctx))
		s.userRepository = repo
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) services.UserService {
	if s.userService == nil {
		service := userService.NewService(s.UserRepository(ctx))
		s.userService = service
	}

	return s.userService
}

func (s *serviceProvider) UserHandlers(ctx context.Context) *handlers.UserV1Handlers {
	if s.userHandlers == nil {
		userHandlers := handlers.NewUserHandlers(s.UserService(ctx))
		s.userHandlers = userHandlers
	}
	return s.userHandlers
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}
