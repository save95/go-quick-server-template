package file

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/save95/go-pkg/http/restful"
	"github.com/save95/go-utils/strutil"
)

type Controller struct {
}

func (c Controller) UploadPublic(ctx *gin.Context) {
	ru := restful.NewResponse(ctx)

	file, _ := ctx.FormFile("file")

	in := uploadRequest{
		Genre:    ctx.Param("genre"),
		Business: ctx.Param("business"),
		File:     file,
	}

	url, err := (service{}).UploadPublic(ctx, &in)
	if nil != err {
		ru.WithError(err)
		return
	}

	ru.Retrieve(map[string]string{
		"url": url,
	})
}

func (c Controller) UploadPublicBase64(ctx *gin.Context) {
	ru := restful.NewResponse(ctx)

	var in base64Request
	if err := ctx.ShouldBindJSON(&in); nil != err {
		ru.WithError(err)
		return
	}

	in.Genre = ctx.Param("genre")
	in.Business = ctx.Param("business")
	url, err := (service{}).UploadPublicBase64(ctx, &in)
	if nil != err {
		ru.WithError(err)
		return
	}

	ru.Retrieve(map[string]string{
		"url": url,
	})
}

func (c Controller) UploadPublicChunk(ctx *gin.Context) {
	ru := restful.NewResponse(ctx)

	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		ru.WithError(errors.New("分块文件太大"))
		return
	}

	file, _ := ctx.FormFile("file")

	in := chunkRequest{
		Genre:       ctx.Param("genre"),
		Business:    ctx.Param("business"),
		Params:      nil,
		File:        file,
		FileId:      strutil.ToInt(ctx.GetHeader("uploader-file-id")),
		ChunksTotal: strutil.ToInt(ctx.GetHeader("uploader-chunks-total")),
		ChunkNumber: strutil.ToInt(ctx.GetHeader("uploader-chunk-number")),
	}

	res, err := (service{}).UploadPublicChunk(ctx, &in)
	if nil != err {
		ru.WithError(err)
		return
	}

	ru.Retrieve(res)
}
