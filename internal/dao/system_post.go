// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"devinggo/internal/dao/internal"
)

// internalSystemPostDao is internal type for wrapping internal DAO implements.
type internalSystemPostDao = *internal.SystemPostDao

// systemPostDao is the data access object for table system_post.
// You can define custom methods on it to extend its functionality as you wish.
type systemPostDao struct {
	internalSystemPostDao
}

var (
	// SystemPost is globally public accessible object for table system_post operations.
	SystemPost = systemPostDao{
		internal.NewSystemPostDao(),
	}
)

// Fill with you ideas below.
