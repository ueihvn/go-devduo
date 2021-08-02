package model

type Field struct {
	ID   uint   `gorm:"primaryKey;auto_increment"`
	Name string `gorm:"not null;unique"`
}

type FieldRepository interface {
	Create(*Field) error
	Get(string) (*Field, error)
	GetAll() ([]Field, error)
	Delete(*Field) error
}
