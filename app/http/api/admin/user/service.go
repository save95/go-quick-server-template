package user

import (
	"context"

	"server-api/global/ecode"
	"server-api/repository/platform"
	"server-api/repository/platform/dao"

	"github.com/save95/go-pkg/model/pager"
	"github.com/save95/go-utils/userutil"
	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"github.com/zywaited/xcopy"
)

type service struct {
}

func (s service) Paginate(_ context.Context, in *paginateRequest) ([]*entity, uint, error) {
	records, total, err := dao.NewUser().Paginate(pager.Option{
		Start: int(in.Start),
		Limit: int(in.GetLimit()),
		Filter: pager.Filter{
			"account": in.Account,
		},
		Sorters: pager.ParseSorts(in.Sort),
	})
	if nil != err {
		return nil, 0, err
	}

	var res []*entity
	if err := xcopy.Copy(&res, records); nil != err {
		return nil, 0, xerror.WrapWithXCode(err, ecode.ErrorVOConverted)
	}

	return res, total, nil
}

func (s service) Create(_ context.Context, in *createRequest) (*entity, error) {
	if err := in.Validate(); nil != err {
		return nil, xerror.WrapWithXCodeStatus(err, xcode.RequestParamError)
	}

	// 判断重复
	_, err := dao.NewUser().FirstByAccount(in.Genre, in.Account)
	if nil == err || !xerror.IsXCode(err, xcode.DBRecordNotFound) {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorRecordExist)
	}

	pwd, err := userutil.NewHasher().Sum(in.Password)
	if nil != err {
		return nil, xerror.Wrap(err, "生成密码失败")
	}

	record := platform.User{
		Genre:     in.Genre,
		Account:   in.Account,
		AvatarURL: in.Avatar,
		Password:  pwd,
		State:     1,
	}
	if err := dao.NewUser().Create(&record); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorSavedData)
	}

	var res entity
	if err := xcopy.Copy(&res, record); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorVOConverted)
	}

	return &res, nil
}

func (s service) Modify(_ context.Context, id uint, in *modifyRequest) (*entity, error) {
	if id == 0 {
		return nil, xerror.WithXCode(ecode.ErrorBadRequest)
	}
	if err := in.Validate(); nil != err {
		return nil, xerror.WrapWithXCodeStatus(err, xcode.RequestParamError)
	}

	record, err := dao.NewUser().First(id)
	if nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorRequestData)
	}

	// 修改名称
	if record.Account != in.Account {
		// 判断重复
		_, err := dao.NewUser().FirstByAccount(record.Genre, in.Account)
		if nil == err || !xerror.IsXCode(err, xcode.DBRecordNotFound) {
			return nil, xerror.WrapWithXCode(err, ecode.ErrorRecordExist)
		}

		record.Account = in.Account
	}

	record.AvatarURL = in.AvatarURL
	record.State = in.State

	// 如果密码不为空，则修改密码
	if len(in.Password) != 0 {
		pwd, err := userutil.NewHasher().Sum(in.Password)
		if nil != err {
			return nil, xerror.Wrap(err, "生成密码失败")
		}
		record.Password = pwd
	}

	if err := dao.NewUser().Save(record); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorSavedData)
	}

	var res entity
	if err := xcopy.Copy(&res, record); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorVOConverted)
	}

	return &res, nil
}
