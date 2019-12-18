package parser

import (
	"favorite-jobs/cache"
	"favorite-jobs/crawler"
	"favorite-jobs/log"
	"favorite-jobs/types"
	"favorite-jobs/utils"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

var err error

func JobIndex(i int, s *goquery.Selection) {
	job := types.Job{PlatformId: 1}

	// 给结构体赋值
	job.Name, _ = s.Attr("data-positionname")
	companyId, _ := s.Attr("data-companyid")
	job.PlatformCompanyId, _ = strconv.ParseInt(companyId, 10, 64)
	platformJobId, _ := s.Attr("data-positionid")
	job.PlatformJobId, _ = strconv.ParseInt(platformJobId, 10, 64)
	job.Salary, _ = s.Attr("data-salary")
	job.MinSalary, job.MaxSalary, err = utils.CalculateSalaryRange(job.Salary)
	utils.CheckErr(err, "计算薪资区间失败")
	job.DetailLink, _ = s.Find(".position_link").Attr("href")
	// 将抓取到的不完整数据先存起来
	cache.JobSave[job.PlatformJobId] = &job
	// 剩下的信息需要去 detail 页面抓取
	crawler.Visit(job.DetailLink, "body", jobDetailAndCompany)
}

// 获取公司信息有两种方案
// 1. 访问公司详情页抓取内容(更全面)
// 2. 直接通过职位详情页右侧抓取(少抓取一次页面 信息较少)

func jobDetailAndCompany(i int, s *goquery.Selection) {
	companyId, _:= s.Find(".target_position").Attr("value")
	// 将获取到的公司Id转成 int64
	cId, _ := strconv.ParseInt(companyId, 10, 64)

	job, ok := cache.JobSave[cId]
	if !ok {
		log.ZapLog.Infow("未找到对应的 job 信息", "c_id:", cId)
		return
	}

	// job
	jobDetail(job, s)

	// company
	company := types.Company{
		PlatformId:  1,
	}
	companyInfo(&company, s)
}

func jobDetail(job *types.Job, s *goquery.Selection) {
	job.Weal = strings.TrimSpace(s.Find(".job-advantage p").Text())
	job.Detail, _ = s.Find(".job-detail").Html()
	strings.TrimSpace(job.Detail)
	s.Find(".job_request h3 span").Each(func(j int, se *goquery.Selection) {
		res := se.Text()
		res = strings.TrimSpace(strings.ReplaceAll(res, "/", ""))
		switch j {
		case 1:
			job.Location = res
		case 2:
			job.Experience = res
		case 3:
			job.Education = res
		}
	})
	// 推送进 channel
	cache.JobChan <- job
}