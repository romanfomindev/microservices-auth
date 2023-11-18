package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/romanfomindev/microservices-auth/internal/config"
	"github.com/romanfomindev/microservices-auth/internal/config/env"
	accessHandlers "github.com/romanfomindev/microservices-auth/internal/handlers/access_v1"
	authHandlers "github.com/romanfomindev/microservices-auth/internal/handlers/auth_v1"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/repositories/url_protected"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user"
	"github.com/romanfomindev/microservices-auth/internal/services"
	accessService "github.com/romanfomindev/microservices-auth/internal/services/access"
	authService "github.com/romanfomindev/microservices-auth/internal/services/auth"
	userService "github.com/romanfomindev/microservices-auth/internal/services/user"
	"github.com/romanfomindev/platform_common/pkg/closer"
	"github.com/romanfomindev/platform_common/pkg/db"
	"github.com/romanfomindev/platform_common/pkg/db/pg"
	"github.com/romanfomindev/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgPool                  *pgxpool.Pool
	dbClient                db.Client
	appConfig               config.AppConfig
	pgConfig                config.PGConfig
	swaggerConfig           config.SwaggerConfig
	authConfig              config.AuthConfig
	loggerConfig            config.LoggerConfig
	txManager               db.TxManager
	grpcConfig              config.GRPCConfig
	httpConfig              config.HTTPConfig
	prometheusConfig        config.PrometheusConfig
	userRepository          repositories.UserRepository
	urlsProtectedRepository repositories.UrlsProtectedRepository
	userService             services.UserService
	authService             services.AuthService
	accessService           services.AccessService
	userHandlers            *handlers.UserV1Handlers
	authHandlers            *authHandlers.AuthV1Handlers
	accessHandlers          *accessHandlers.AccessV1Handlers
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

func (s *serviceProvider) AppConfig() config.AppConfig {
	if s.appConfig == nil {
		cfg, err := env.NewAppConfig()
		if err != nil {
			log.Fatalf("failed to get app config: %s", err.Error())
		}

		s.appConfig = cfg
	}

	return s.appConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
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

func (s *serviceProvider) HTTPConfig() config.GRPCConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) AuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
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
		s.userRepository = user.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UrlsProtectedRepository(ctx context.Context) repositories.UrlsProtectedRepository {
	if s.urlsProtectedRepository == nil {
		s.urlsProtectedRepository = url_protected.NewUrlProtectedPepository(s.DBClient(ctx))
	}

	return s.urlsProtectedRepository
}

func (s *serviceProvider) UserService(ctx context.Context) services.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) services.AuthService {
	if s.authService == nil {
		authConfig := s.AuthConfig()
		s.authService = authService.NewAuthService(
			authConfig.RefreshTokenSecretKey(),
			authConfig.AccessTokenSecretKey(),
			authConfig.RefreshTokenExpiration(),
			authConfig.AccessTokenExpiration(),
			s.UserRepository(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) services.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewAccessService(s.UrlsProtectedRepository(ctx), s.AuthConfig().AccessTokenSecretKey())
	}

	return s.accessService
}

func (s *serviceProvider) UserHandlers(ctx context.Context) *handlers.UserV1Handlers {
	if s.userHandlers == nil {
		s.userHandlers = handlers.NewUserHandlers(s.UserService(ctx))
	}
	return s.userHandlers
}

func (s *serviceProvider) AuthHandlers(ctx context.Context) *authHandlers.AuthV1Handlers {
	if s.authHandlers == nil {
		s.authHandlers = authHandlers.NewAuthHandlers(s.AuthService(ctx))
	}
	return s.authHandlers
}

func (s *serviceProvider) AccessHandlers(ctx context.Context) *accessHandlers.AccessV1Handlers {
	if s.accessHandlers == nil {
		s.accessHandlers = accessHandlers.NewAccessHandlers(s.AccessService(ctx))
	}
	return s.accessHandlers
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN(), s.PGConfig().Timeout())
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
