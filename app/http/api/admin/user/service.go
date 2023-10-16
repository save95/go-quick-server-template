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
	records, total, err := dao.NewVWUser().Paginate(pager.Option{
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
	_, err := dao.NewVWUser().FirstByAccount(in.Account)
	if nil == err || !xerror.IsXCode(err, xcode.DBRecordNotFound) {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorRecordExist)
	}

	pwd, err := userutil.NewHasher().Sum(in.Password)
	if nil != err {
		return nil, xerror.Wrap(err, "生成密码失败")
	}

	record := platform.User{
		Account:   in.Account,
		AvatarURL: in.Avatar,
		Password:  pwd,
		State:     1,
	}
	stat := platform.UserStat{}
	if err := dao.NewUser().Create(in.GetGenres(), &record, &stat); nil != err {
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

	record, err := dao.NewVWUser().First(id, "UserRoles")
	if nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorRequestData)
	}

	roles := make([]int8, 0)
	for _, role := range record.UserRoles {
		roles = append(roles, int8(role.Genre))
	}

	// 修改名称
	if record.Account != in.Account {
		// 判断重复
		_, err := dao.NewVWUser().FirstByAccount(in.Account)
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

	if err := dao.NewUser().Update(record.ToUser(), roles); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorSavedData)
	}

	var res entity
	if err := xcopy.Copy(&res, record); nil != err {
		return nil, xerror.WrapWithXCode(err, ecode.ErrorVOConverted)
	}

	return &res, nil
}
