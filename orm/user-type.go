package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type UserType struct { // 01 Seller 02 Customer
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(100)"`
}
