// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
)

type CompanyTeamMember = internal.CompanyTeamMemberDao

func NewCompanyTeamMember(dao ...co_interface.IDao) *CompanyTeamMember {
	return internal.NewCompanyTeamMemberDao(dao...)
}
