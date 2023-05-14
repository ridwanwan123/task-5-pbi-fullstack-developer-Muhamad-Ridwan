package models

type User struct {
	Id       			int64   `gorm:"primaryKey" json:"id"`
	Username		 	string  `gorm:"type:varchar(150)" json:"username"`
	Email				 	string  `gorm:"type:varchar(150)" json:"email"`
	Password		 	string  `gorm:"type:varchar(150)" json:"password"`
}