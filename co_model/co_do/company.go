// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Company is the golang structure of table co_company for DAO operations like Where/Data.
type Company struct {
	g.Meta        `orm:"table:co_company, do:true"`
	Id            interface{} // ID
	Name          interface{} // 名称
	ContactName   interface{} // 商务联系人
	ContactMobile interface{} // 商务联系电话
	UserId        interface{} // 管理员ID
	ParentId      interface{} // 父级ID
	State         interface{} // 状态：0未启用，1正常
	Remark        interface{} // 备注
	CreatedBy     interface{} // 创建者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedBy     interface{} // 更新者
	UpdatedAt     *gtime.Time // 更新时间
	DeletedBy     interface{} // 删除者
	DeletedAt     *gtime.Time // 删除时间
}
