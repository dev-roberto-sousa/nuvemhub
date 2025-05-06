package model

type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	UserName string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);not null"`
	Status   bool   `gorm:"default:true"`
}
