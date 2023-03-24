package model

type User struct {
	ID           int64  `gorm:"primaryKey;not_null;auto_increment"`
	UserName     string `gorm:"unique_index;not_null"`
	FirstName    string
	HashPassword string
}
