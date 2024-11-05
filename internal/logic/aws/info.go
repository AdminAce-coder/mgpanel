package aws

import (
	"context"
	"mgpanel/internal/dao"
	"mgpanel/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

type sGetinfo struct{}

func init() {
	service.RegisterGetinfo(New())
}
func New() *sGetinfo {
	return &sGetinfo{}
}

// 获取所有实例信息
func (s *sGetinfo) GetAll(ctx context.Context) string {

	db := dao.InInfo.DB()
	result, err := db.Model("In_info").All()
	if err != nil {
		db.GetLogger().Error(ctx, err.Error())
	}
	return result.Json()
}
