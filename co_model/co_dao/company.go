// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
	"github.com/SupenBysz/gf-admin-company-modules/utility/dao_helper"
)

type CompanyDao = internal.CompanyDao

var (
	// Company is globally public accessible object for table pro_company operations.
	Company = func(module co_interface.IModules) dao_helper.IDao[internal.CompanyColumns] {
		return dao_helper.NewDao[internal.CompanyColumns](module.GetConfig(), internal.NewCompanyDao())
	}
)
