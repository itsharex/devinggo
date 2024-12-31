// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemMenuDao is internal type for wrapping internal DAO implements.
type internalSystemMenuDao = *internal.SystemMenuDao

// systemMenuDao is the data access object for table system_menu.
// You can define custom methods on it to extend its functionality as you wish.
type systemMenuDao struct {
	internalSystemMenuDao
}

var (
	// SystemMenu is globally public accessible object for table system_menu operations.
	SystemMenu = systemMenuDao{
		internal.NewSystemMenuDao(),
	}
)

// Fill with you ideas below.
