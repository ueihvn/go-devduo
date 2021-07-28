package model

type Field struct {
	ID   uint   `gorm:"primaryKey;auto_increment"`
	Name string `gorm:"not null;unique"`
}
