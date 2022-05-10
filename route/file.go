package route

import (
	"server-api/app/http/api/file"
	"server-api/global"

	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/middleware"
	"github.com/save95/go-pkg/http/types"
)

func registerFile(rg *gin.Engine) {
	api := file.Controller{}

	v1 := rg.Group(
		"/file",
		middleware.RESTFul(global.ApiVersionLatest),
		middleware.JWTWith(global.JWTOption(false)),
		middleware.Roles([]types.IRole{global.RoleUser}),
	)
	{
		v1.POST("/by-simple/:genre/:business", api.UploadPublic)

		v1.POST("/by-base64/:genre/:business", api.UploadPublicBase64)

		// 配合 npm install huge-uploader --save 食用
		// https://www.npmjs.com/package/huge-uploader
		v1.POST("/by-chunk/:genre/:business", api.UploadPublicChunk)
	}
}
