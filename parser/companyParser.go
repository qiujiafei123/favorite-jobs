package parser

import (
	"favorite-jobs/cache"
	"favorite-jobs/types"
	"github.com/PuerkitoBio/goquery"
)

func companyInfo(company *types.Company, s *goquery.Selection) {
	company.CompanyUrl, _ = s.Find("#job_company a").Attr("href")
	company.CompanyLogo, _ = s.Find("#job_company b2").Attr("src")
	company.CompanyLogo = "http://" + company.CompanyLogo
	company.CompanyName, _ = s.Find("#job_company b2").Attr("alt")
	s.Find("#job_company .c_feature_name").Each(func(i int, selection *goquery.Selection) {
		switch i {
		case 0:
			company.Industry = selection.Text()
		case 1:
			company.Growth = selection.Text()
		case 2:
			company.People = selection.Text()
		case 3:
			company.CompanyOfficeUrl = selection.Text()
		}
	})
	// 推送进 channel
	cache.CompanyChan <- company
}
