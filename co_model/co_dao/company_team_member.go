// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
)

type CompanyTeamMember = internal.CompanyTeamMemberDao

func NewCompanyTeamMember[T co_interface.IDao](dao T) T {
	var result interface{} = internal.NewCompanyTeamMemberDao(dao)

	if v, ok := result.(T); ok {
		return v
	}

	return dao
}
