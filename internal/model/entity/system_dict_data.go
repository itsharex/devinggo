// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictData is the golang structure for table system_dict_data.
type SystemDictData struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	TypeId    int64       `json:"typeId"    orm:"type_id"    description:""` //
	Label     string      `json:"label"     orm:"label"      description:""` //
	Value     string      `json:"value"     orm:"value"      description:""` //
	Code      string      `json:"code"      orm:"code"       description:""` //
	Sort      int         `json:"sort"      orm:"sort"       description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""` //
	Remark    string      `json:"remark"    orm:"remark"     description:""` //
}
