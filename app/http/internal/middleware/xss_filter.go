package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/xss"
)

func XSSFilter() gin.HandlerFunc {
	return middleware.XSSFilter(
		//middleware.WithXSSDebug(),
		middleware.WithXSSGlobalPolicy(xss.PolicyStrict),
		middleware.WithXSSGlobalSkipFields("password"),
		middleware.WithXSSRoutePolicy("admin", xss.PolicyUGC),
		middleware.WithXSSRoutePolicy("/callback/", xss.PolicyNone),
		middleware.WithXSSRoutePolicy("/endpoint", xss.PolicyNone),
		middleware.WithXSSRoutePolicy("/ping", xss.PolicyNone),
	)
}
