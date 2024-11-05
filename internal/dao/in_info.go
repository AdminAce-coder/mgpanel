// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mgpanel/internal/dao/internal"
)

// internalInInfoDao is internal type for wrapping internal DAO implements.
type internalInInfoDao = *internal.InInfoDao

// inInfoDao is the data access object for table In_info.
// You can define custom methods on it to extend its functionality as you wish.
type inInfoDao struct {
	internalInInfoDao
}

var (
	// InInfo is globally public accessible object for table In_info operations.
	InInfo = inInfoDao{
		internal.NewInInfoDao(),
	}
)

// Fill with you ideas below.
