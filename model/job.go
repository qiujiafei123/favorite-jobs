package model

import "favorite-jobs/types"

func AddJob(jobFull *types.Job) {
	index := types.Job{PlatformId:jobFull.PlatformId, PlatformJobId:jobFull.PlatformJobId}
	jobId := GetJobRecord(index)
	if jobId == 0 {
		DB.Create(jobFull)
		return
	}
	jobFull.ID = jobId
	DB.Model(jobFull).Updates(*jobFull)
}

func GetJobRecord(jobFull types.Job) int64 {
	DB.Where("platform_job_id = ?", jobFull.PlatformJobId).
		Where("platform_id = ?", jobFull.PlatformId).
		Find(&jobFull)
	if jobFull.ID == 0 {
		return 0
	}
	return jobFull.ID
}