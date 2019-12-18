package types

type Job struct {
	ID int64
	PlatformId int8
	PlatformJobId int64
	PlatformCompanyId int64
	Name string
	Salary string
	MinSalary string
	MaxSalary string
	DetailLink string
	Detail string
	Tag string
	Location string
	Experience string
	Education string
	Weal string
}

func (j Job) TableName() string {
	return "job"
}