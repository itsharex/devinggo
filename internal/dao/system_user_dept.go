// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemUserDeptDao is internal type for wrapping internal DAO implements.
type internalSystemUserDeptDao = *internal.SystemUserDeptDao

// systemUserDeptDao is the data access object for table system_user_dept.
// You can define custom methods on it to extend its functionality as you wish.
type systemUserDeptDao struct {
	internalSystemUserDeptDao
}

var (
	// SystemUserDept is globally public accessible object for table system_user_dept operations.
	SystemUserDept = systemUserDeptDao{
		internal.NewSystemUserDeptDao(),
	}
)

// Fill with you ideas below.
