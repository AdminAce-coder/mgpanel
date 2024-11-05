// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// InInfo is the golang structure for table In_info.
type InInfo struct {
	Region       string      `json:"region"       orm:"region"        description:"区域"`
	Iname        string      `json:"iname"        orm:"iname"         description:"实例名"`
	Ctime        *gtime.Time `json:"ctime"        orm:"ctime"         description:"创建时间"`
	State        string      `json:"state"        orm:"state"         description:"实例状态"`
	Itype        string      `json:"itype"        orm:"itype"         description:"实例类型"`
	Ip           string      `json:"ip"           orm:"ip"            description:""`
	Group        string      `json:"group"        orm:"group"         description:""`
	Stime        string      `json:"stime"        orm:"stime"         description:"实例剩余时间"`
	ResourceType string      `json:"resourceType" orm:"Resource_type" description:"资源类型"`
	Tag          string      `json:"tag"          orm:"Tag"           description:"实例标签  实例标签"`
}
