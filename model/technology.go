package model

type Technology struct {
	ID   uint   `gorm:"primaryKey;auto_increment" json:"id,omitempty"`
	Name string `gorm:"not null;unique" json:"name,omitempty"`
}

type TechnologyRepository interface {
	Create(*Technology) error
	Get(string) (*Technology, error)
	GetAll() ([]Technology, error)
	Delete(*Technology) error
}
