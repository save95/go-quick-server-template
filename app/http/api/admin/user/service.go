package user

import (
	"context"

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
		Sorters: nil,
	})
	if nil != err {
		return nil, 0, err
	}

	var res []*entity
	if err := xcopy.Copy(&res, records); nil != err {
		return nil, 0, xerror.Wrap(err, "data convert error")
	}

	return res, total, nil
}

func (s service) Create(_ context.Context, in *createRequest) (*entity, error) {
	if err := in.Validate(); nil != err {
		return nil, xerror.WithXCodeMessage(xcode.RequestParamError, err.Error())
	}

	// 判断重复
	udao := dao.NewUser()
	_, err := udao.FirstByAccount(in.Genre, in.Account)
	if nil == err || !xerror.IsXCode(err, xcode.DBRecordNotFound) {
		return nil, xerror.Wrap(err, "数据已存在")
	}

	pwd, err := userutil.NewHasher().Sum(in.Password)
	if nil != err {
		return nil, xerror.Wrap(err, "生成密码失败")
	}

	record := platform.User{
		Genre:    in.Genre,
		Account:  in.Account,
		Avatar:   in.Avatar,
		Password: pwd,
		State:    1,
	}
	if err := udao.Create(&record); nil != err {
		return nil, xerror.Wrap(err, "新增失败")
	}

	var res entity
	if err := xcopy.Copy(&res, record); nil != err {
		return nil, xerror.Wrap(err, "数据转换失败")
	}

	return &res, nil
}

func (s service) Modify(_ context.Context, id uint, in *modifyRequest) (*entity, error) {
	idao := dao.NewUser()
	record, err := idao.First(id)
	if nil != err {
		return nil, xerror.Wrap(err, "数据错误")
	}

	// 修改名称
	if record.Account != in.Account {
		// 判断重复
		_, err := idao.FirstByAccount(record.Genre, in.Account)
		if nil == err || !xerror.IsXCode(err, xcode.DBRecordNotFound) {
			return nil, xerror.Wrap(err, "数据已存在")
		}

		record.Account = in.Account
	}

	record.Avatar = in.Avatar
	record.State = in.State

	// 如果密码不为空，则修改密码
	if len(in.Password) != 0 {
		pwd, err := userutil.NewHasher().Sum(in.Password)
		if nil != err {
			return nil, xerror.Wrap(err, "生成密码失败")
		}
		record.Password = pwd
	}

	if err := idao.Save(record); nil != err {
		return nil, xerror.Wrap(err, "新增失败")
	}

	var res entity
	if err := xcopy.Copy(&res, record); nil != err {
		return nil, xerror.Wrap(err, "数据转换失败")
	}

	return &res, nil
}
