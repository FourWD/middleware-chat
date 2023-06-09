package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type User struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	DomainID string `db:"domain_id" json:"domain_id" gorm:"type:varchar(36);"`

	AppID      string `db:"app_id" json:"app_id" gorm:"type:varchar(10);"`
	SaleID     string `db:"sale_id" json:"sale_id" gorm:"type:varchar(36);"`
	StatusID   string `db:"status_id" json:"status_id" gorm:"type:varchar(36);"`
	UserTypeID string `db:"UserTypeID" json:"UserTypeID" gorm:"type:varchar(2);"`

	Username  string `db:"username" json:"username" gorm:"type:varchar(50)"`
	Password  string `db:"password" json:"password" gorm:"type:varchar(100)"`
	Firstname string `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname  string `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Nickname  string `db:"nickname" json:"nickname" gorm:"type:varchar(50)"`
	Birthday  string `db:"birthday" json:"birthday" `
	Avatar    string `db:"avatar" json:"avatar" gorm:"type:varchar(100)"`
	Email     string `db:"email" json:"email" gorm:"type:varchar(50)"`
	Position  string `db:"position" json:"position" gorm:"type:varchar(100)"`
	NotiToken string `db:"noti_token" json:"noti_token" gorm:"type:varchar(50);"`
	RowOrder  int    `db:"row_order" json:"row_order" gorm:"type:int"`
	IsQueue   bool   `db:"is_queue" json:"is_queue" gorm:"type:varchar(1)"`
}
