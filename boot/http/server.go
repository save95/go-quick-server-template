package http

import (
	"context"
	"log"
	"net/http"
	m "server-api/app/http/middleware"
	"server-api/global"
	"server-api/route"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/restful"
	"github.com/save95/xlog"
)

type server struct {
	ctx        context.Context
	httpServer *http.Server
}

func NewHttpServer(ctx context.Context) *server {
	return &server{ctx: ctx}
}

func (s *server) Start() error {
	// 验证器使用 validator.v9
	//binding.Validator = new(validator.DefaultValidator)

	r := gin.New()

	// 注册全局中间件。注意顺序不要随意调整
	r.Use(gin.Recovery())
	r.Use(m.CORS())
	r.Use(middleware.HttpContext())
	//r.Use(middleware.RESTFul(global.ApiVersionLatest))
	//r.Use(middleware.Log())
	//r.Use(middleware.ParserSession())

	// 开启 http 日志
	if global.Config.Log.HttpLog && global.Log != nil {
		r.Use(middleware.HttpLogger(restful.LogOption{
			Logger:    global.Log,
			OnlyError: global.Config.Log.HttpLogOnlyError,
		}))
	}

	// 非调试模式下，启用发布模式
	if xlog.ParseLevel(global.Config.Log.Level) != xlog.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册 swagger
	registerSwagger(r)

	// 注册路由，路由统一安置在 app/route 目录，由 main 引导
	route.Register(r)

	global.Log.Infof("http server Listening and serving HTTP on %s", global.Config.Server.Addr)
	global.Log.Info("http server starting...")
	//err := r.Run(global.Config.Server.Addr)

	s.httpServer = &http.Server{
		Addr:    global.Config.Server.Addr,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("httpServer start failed: %s\n", err)
		}
	}()

	return nil
}

func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
