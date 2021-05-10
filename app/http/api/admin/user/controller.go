package user

import (
	"github.com/gin-gonic/gin"
	"github.com/save95/go-pkg/http/restful"
	"github.com/save95/go-pkg/utils/strutil"
)

type Controller struct {
}

func (c Controller) Paginate(ctx *gin.Context) {
	rru := restful.NewResponse(ctx)

	var in paginateRequest
	if err := ctx.ShouldBindQuery(&in); nil != err {
		rru.WithError(err)
		return
	}

	records, total, err := service{}.Paginate(ctx, &in)
	if nil != err {
		rru.WithError(err)
		return
	}

	rru.ListWithPagination(total, records)
}

func (c Controller) Create(ctx *gin.Context) {
	rru := restful.NewResponse(ctx)

	var in createRequest
	if err := ctx.ShouldBindJSON(&in); nil != err {
		rru.WithError(err)
		return
	}

	record, err := service{}.Create(ctx, &in)
	if nil != err {
		rru.WithError(err)
		return
	}

	rru.Post(record)
}

func (c Controller) Modify(ctx *gin.Context) {
	rru := restful.NewResponse(ctx)

	id := strutil.ToIntWith(ctx.Param("id"), 0)

	var in modifyRequest
	if err := ctx.ShouldBindJSON(&in); nil != err {
		rru.WithError(err)
		return
	}

	record, err := service{}.Modify(ctx, uint(id), &in)
	if nil != err {
		rru.WithError(err)
		return
	}

	rru.Post(record)
}
