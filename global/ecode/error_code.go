package ecode

import "github.com/save95/xerror/xcode"

var (
	ErrorBadRequest  = xcode.NewWithMessage(3001, "请求数据错误或不存在")
	ErrorVOConverted = xcode.NewWithMessage(3002, "数据转换失败")
	ErrorRequestData = xcode.NewWithMessage(3003, "请求数据错误")
	ErrorSavedData   = xcode.NewWithMessage(3004, "数据保存失败")
	ErrorRecordExist = xcode.NewWithMessage(3005, "数据已存在")

	ErrorAuthParams = xcode.NewWithMessage(4000, "账号或密码错误")
	ErrorAuthFailed = xcode.NewWithMessage(4001, "授权登录失败")

	// todo 其他业务错误码
)
