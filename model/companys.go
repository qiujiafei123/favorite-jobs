package model

import "favorite-jobs/types"

func AddCompany(company *types.Company) {
	companyId := GetCompanyRecord(types.Company{CompanyId:company.CompanyId, PlatformId:company.PlatformId})
	if companyId == 0 {
		DB.Create(company)
		return
	}
	company.Id = companyId
	DB.Model(company).Updates(*company)
}

func GetCompanyRecord(company types.Company) int64 {
	DB.Where("company_id = ?", company.CompanyId).
		Where("platform_id = ?", company.PlatformId).
		Find(&company)
	if company.Id == 0 {
		return 0
	}
	return company.Id
}