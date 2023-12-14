package Models

type Course struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null" validate:"required"`
	Description string  `json:"description" gorm:"not null" validate:"required"`
	Price       float64 `json:"price" gorm:"not null" validate:"required,gt=0"`
	Mentor      string  `json:"mentor" gorm:"not null" validate:"required"`
}

// returning rtable name
func (course *Course) TableName() string {
	return "course"
}
