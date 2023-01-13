// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
	"github.com/SupenBysz/gf-admin-company-modules/utility/dao_helper"
)

type CompanyTeamMemberDao = internal.CompanyTeamMemberDao

var (
	// CompanyTeamMember is globally public accessible object for table pro_company_team_member operations.
	CompanyTeamMember = func(module co_interface.IModules) dao_helper.IDao[internal.CompanyTeamMemberColumns] {
		return dao_helper.NewDao[internal.CompanyTeamMemberColumns](module.GetConfig(), internal.NewCompanyTeamMemberDao())
	}
)
