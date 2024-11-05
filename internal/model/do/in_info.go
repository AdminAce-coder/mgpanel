// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// InInfo is the golang structure of table In_info for DAO operations like Where/Data.
type InInfo struct {
	g.Meta       `orm:"table:In_info, do:true"`
	Region       interface{} // 区域
	Iname        interface{} // 实例名
	Ctime        *gtime.Time // 创建时间
	State        interface{} // 实例状态
	Itype        interface{} // 实例类型
	Ip           interface{} //
	Group        interface{} //
	Stime        interface{} // 实例剩余时间
	ResourceType interface{} // 资源类型
	Tag          interface{} // 实例标签  实例标签
}
