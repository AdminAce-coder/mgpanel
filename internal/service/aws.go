// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGetinfo interface {
		// 获取所有实例信息
		GetAll(ctx context.Context) string
	}
)

var (
	localGetinfo IGetinfo
)

func Getinfo() IGetinfo {
	if localGetinfo == nil {
		panic("implement not found for interface IGetinfo, forgot register?")
	}
	return localGetinfo
}

func RegisterGetinfo(i IGetinfo) {
	localGetinfo = i
}
