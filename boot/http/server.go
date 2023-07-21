package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"server-api/global"
	"server-api/route"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/save95/go-pkg/application"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/xss"
	"github.com/save95/xlog"
)

type server struct {
	ctx        context.Context
	httpServer *http.Server
}

func NewHttpServer(ctx context.Context) application.IApplication {
	return &server{ctx: ctx}
}

func (s *server) Start() error {
	// 验证器使用 validator.v9
	//binding.Validator = new(validator.DefaultValidator)

	r := gin.New()

	// 注册全局中间件。注意顺序不要随意调整
	r.Use(gin.Recovery())
	r.Use(middleware.CORS(
		middleware.WithCORSAllowOriginFunc(func(origin string) bool {
			if !global.Env().IsProd() {
				return true
			}

			// todo cors domain
			//return origin == "https://xxxx.com"
			return true
		}),
		middleware.WithCORSAllowHeaders("X-Custom-Key"),
		middleware.WithCORSExposeHeaders("X-Custom-Key"),
		middleware.WithCORSMaxAge(24*time.Hour),
	))

	// 开启 http 缓存
	if global.Config.HttpCache.Enabled {
		r.Use(middleware.HttpCache(
			//middleware.WithHttpCacheDebug(),
			middleware.WithHttpCacheLogger(global.Log),
			middleware.WithHttpCacheJWTOption(global.JWTOption(false)),
			middleware.WithHttpCacheGlobalDuration(5*time.Minute),
			middleware.WithHttpCacheRedisStore(redis.NewClient(&redis.Options{
				Addr:     global.Config.HttpCache.Addr,
				Password: global.Config.HttpCache.Password,
				DB:       global.Config.HttpCache.DB,
			})),
			middleware.WithHttpCacheGlobalSkipFields("v"),
			middleware.WithHttpCacheRouteSkipFiledPolicy("/user/", true),
		))
	}

	r.Use(middleware.XSSFilter(
		//middleware.WithXSSDebug(),
		middleware.WithXSSGlobalPolicy(xss.PolicyStrict),
		middleware.WithXSSGlobalSkipFields("password"),
		middleware.WithXSSRoutePolicy("admin", xss.PolicyUGC),
		middleware.WithXSSRoutePolicy("/callback/", xss.PolicyNone),
		middleware.WithXSSRoutePolicy("/endpoint", xss.PolicyNone),
		middleware.WithXSSRoutePolicy("/ping", xss.PolicyNone),
	))
	r.Use(middleware.HttpContext())
	//r.Use(middleware.RESTFul(global.ApiVersionLatest))
	//r.Use(middleware.Log())
	//r.Use(middleware.ParserSession())

	// 开启 http 日志
	if global.Config.Log.HttpLog && global.Log != nil {
		r.Use(middleware.HttpLogger(middleware.HttpLoggerOption{
			Logger:    global.Log,
			OnlyError: global.Config.Log.HttpLogOnlyError,
		}))
	}

	// 非调试模式下，启用发布模式
	if xlog.ParseLevel(global.Config.Log.Level) != xlog.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

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
