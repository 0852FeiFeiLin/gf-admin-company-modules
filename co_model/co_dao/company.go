// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
	"github.com/SupenBysz/gf-admin-company-modules/utility/dao_helper"
)

type internalCompanyDao = *dao_helper.CustomDao[internal.CompanyColumns]

// companyDao 是 company 表的数据访问对象。
// 您可以在其上定义自定义方法，以根据需要扩展其功能。
type companyDao struct {
	internalCompanyDao
}

var (
	// Company 表 pro_company 操作的全局公共可访问对象。
	Company = func(module co_interface.IModules) *companyDao {
		return &companyDao{
			internalCompanyDao: dao_helper.NewDao[internal.CompanyColumns](
				module.GetConfig(),
				&dao_helper.CustomDao[internal.CompanyColumns]{
					IDao: internal.NewCompanyDao(),
				},
			),
		}
	}
)
