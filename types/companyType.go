package types

type Company struct {
	Id int64
	PlatformId int64
	CompanyId int64
	CompanyName string
	Industry string
	Growth string
	People string
	CompanyUrl string
	CompanyLogo string
	CompanyOfficeUrl string
}

func (c Company) TableName() string {
	return "company"
}