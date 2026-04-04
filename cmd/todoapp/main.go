package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_pgx_pool "github.com/Mihail4531/golang-todo/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Mihail4531/golang-todo/internal/core/transport/http/middleware"
	core_http_server "github.com/Mihail4531/golang-todo/internal/core/transport/server"
	tasks_postgres_repository "github.com/Mihail4531/golang-todo/internal/features/tasks/repository/postgres"
	tasks_service "github.com/Mihail4531/golang-todo/internal/features/tasks/service"
	tasks_transport_http "github.com/Mihail4531/golang-todo/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/Mihail4531/golang-todo/internal/features/users/repository/postgres"
	users_service "github.com/Mihail4531/golang-todo/internal/features/users/service"
	users_transport_http "github.com/Mihail4531/golang-todo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {
		time.Local = timeZone


	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()
	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger %w", err)
		os.Exit(1)
	}
	defer logger.Close()
	logger.Debug("initial postgres pool ")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTPP := users_transport_http.NewUsersHTTPHandler(usersService)
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	taskTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)
	logger.Debug("initial HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	ApiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	ApiVersionRouter.RegisterRoutes(usersTransportHTPP.Routes()...)
	ApiVersionRouter.RegisterRoutes(taskTransportHTTP.Routes()...)
	httpServer.RegisterApiRouters(ApiVersionRouter)
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
