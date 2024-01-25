package models

type CustomerCompany struct {
	ID          uint       `gorm:"primarykey"`
	CompanyName string     `gorm:"type:varchar;column:company_name" json:"company_name"`
	Customers   []Customer `json:"customers" gorm:"foreignKey:CompanyID"`
}
