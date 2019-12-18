package cache

import (
	"favorite-jobs/model"
	"favorite-jobs/types"
	"fmt"
	"sync"
)



var (
	group sync.WaitGroup
	JobChan chan *types.Job
	CompanyChan chan *types.Company
)

func InitModelChan() {
	if JobChan == nil {
		JobChan = make(chan *types.Job, 20)
	}

	if CompanyChan == nil {
		CompanyChan = make(chan *types.Company, 20)
	}
}

func Revice(g *sync.WaitGroup) {
	group.Add(2)
	go reciveJob()
	go reciveCompany()
	group.Wait()
	g.Done()
}

func reciveJob() {
	for job := range JobChan {
		// 接收jobs
		model.AddJob(job)
		fmt.Println("收到职位信息")
	}
	defer group.Done()
}

func reciveCompany() {
	for company := range CompanyChan {
		model.AddCompany(company)
		fmt.Println("收到公司信息")
	}
	defer group.Done()
}