// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InInfoDao is the data access object for table In_info.
type InInfoDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns InInfoColumns // columns contains all the column names of Table for convenient usage.
}

// InInfoColumns defines and stores column names for table In_info.
type InInfoColumns struct {
	Region       string // 区域
	Iname        string // 实例名
	Ctime        string // 创建时间
	State        string // 实例状态
	Itype        string // 实例类型
	Ip           string //
	Group        string //
	Stime        string // 实例剩余时间
	ResourceType string // 资源类型
	Tag          string // 实例标签  实例标签
}

// inInfoColumns holds the columns for table In_info.
var inInfoColumns = InInfoColumns{
	Region:       "region",
	Iname:        "iname",
	Ctime:        "ctime",
	State:        "state",
	Itype:        "itype",
	Ip:           "ip",
	Group:        "group",
	Stime:        "stime",
	ResourceType: "Resource_type",
	Tag:          "Tag",
}

// NewInInfoDao creates and returns a new DAO object for table data access.
func NewInInfoDao() *InInfoDao {
	return &InInfoDao{
		group:   "default",
		table:   "In_info",
		columns: inInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InInfoDao) Columns() InInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
