package Models

type Organism struct {
	ID uint64   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}