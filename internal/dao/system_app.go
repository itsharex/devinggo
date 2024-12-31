// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemAppDao is internal type for wrapping internal DAO implements.
type internalSystemAppDao = *internal.SystemAppDao

// systemAppDao is the data access object for table system_app.
// You can define custom methods on it to extend its functionality as you wish.
type systemAppDao struct {
	internalSystemAppDao
}

var (
	// SystemApp is globally public accessible object for table system_app operations.
	SystemApp = systemAppDao{
		internal.NewSystemAppDao(),
	}
)

// Fill with you ideas below.
