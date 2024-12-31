// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemRoleDao is internal type for wrapping internal DAO implements.
type internalSystemRoleDao = *internal.SystemRoleDao

// systemRoleDao is the data access object for table system_role.
// You can define custom methods on it to extend its functionality as you wish.
type systemRoleDao struct {
	internalSystemRoleDao
}

var (
	// SystemRole is globally public accessible object for table system_role operations.
	SystemRole = systemRoleDao{
		internal.NewSystemRoleDao(),
	}
)

// Fill with you ideas below.
