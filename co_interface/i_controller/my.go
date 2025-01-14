package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
)

type IMy interface {
	iModule
	// GetProfile 获取当前员工及用户信息
	GetProfile(ctx context.Context, _ *co_company_api.GetProfileReq) (*co_model.MyProfileRes, error)

	// GetCompany 获取当前公司信息
	GetCompany(ctx context.Context, _ *co_company_api.GetCompanyReq) (*co_model.MyCompanyRes, error)

	// GetTeams 获取当前团队信息
	GetTeams(ctx context.Context, _ *co_company_api.GetTeamsReq) (co_model.MyTeamListRes, error)
}
