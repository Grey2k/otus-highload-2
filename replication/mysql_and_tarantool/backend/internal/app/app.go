package app

import (
	"backend/internal/config"
	"backend/internal/implementation"
	"backend/internal/infrastructure/tarantool"
	httptransport "backend/internal/transport/http"
	wstransport "backend/internal/transport/ws"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const httpTimeoutClose = 5 * time.Second

type Option func(*App)

// WithLogger adding logger option.
func WithLogger(l *zap.Logger) Option {
	return func(a *App) {
		a.logger = l
	}
}

// WithTarantool adding tarantool option.
func WithTarantool(conn *tarantool.Conn) Option {
	return func(a *App) {
		a.tConn = conn
	}
}

// App is main application instance.
type App struct {
	cfg     *config.Config
	httpSrv *http.Server
	tConn   *tarantool.Conn
	wsConns *wstransport.Conns
	logger  *zap.Logger
}

// NewApp returns instance of app.
func NewApp(cfg *config.Config, opts ...Option) *App {
	app := &App{
		cfg:     cfg,
		wsConns: wstransport.NewWSConnects(),
		logger:  zap.NewNop(),
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Run lunch application.
func (a *App) Run(mysqlConn *sql.DB) {
	authSvc := implementation.NewAuthService(implementation.NewUserRepository(mysqlConn), a.cfg.JWT)
	socialSvc := implementation.NewSocialService(implementation.NewUserRepository(mysqlConn))
	//socialSvc := implementation.NewSocialService(implementation.NewTarantoolRepository(a.tConn))
	messengerSvc := implementation.NewMessengerService(
		implementation.NewUserRepository(mysqlConn),
		implementation.NewMessengerRepository(mysqlConn))

	a.httpSrv = httptransport.NewHTTPServer(a.cfg.Addr, httptransport.MakeEndpoints(authSvc, socialSvc, messengerSvc))

	go func() {
		if err := a.httpSrv.ListenAndServe(); err != nil {
			a.logger.Fatal("http serve error, ", zap.Error(err))
		}
	}()
}

func (a *App) Stop() {
	a.wsConns.Close()

	ctx, cancel := context.WithTimeout(context.Background(), httpTimeoutClose)
	defer cancel()

	if err := a.httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("http closing error, ", zap.Error(err))
	}
}
