package Models

type Recipient struct {
	Name string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}
