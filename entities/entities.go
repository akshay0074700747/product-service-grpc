package entities

type Products struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price int
	Stock int
}
