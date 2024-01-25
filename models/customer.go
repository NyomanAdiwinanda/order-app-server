package models

type Customer struct {
	ID          uint            `gorm:"primarykey"`
	IdName      string          `gorm:"type:varchar;column:id_name" json:"id_name"`
	Login       string          `gorm:"type:varchar;column:login" json:"login"`
	Password    string          `gorm:"type:varchar;column:password" json:"password"`
	Name        string          `gorm:"type:varchar;column:name" json:"name"`
	CompanyID   uint            `gorm:"column:company_id" json:"company_id"`
	CreditCards string          `gorm:"type:text;column:credit_cards" json:"credit_cards"`
	Orders      []Order         `json:"orders" gorm:"foreignKey:CustomerID"`
	Company     CustomerCompany `json:"company" gorm:"foreignKey:CompanyID"`
}
