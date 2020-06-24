package models

type RecipientsList struct {
	ID uint64 `gorm: "primary_key"`
	Name  string `gorm:"size:255"`
	Recipients []*Recipient `gorm:"many2many:mail_lists;"`
	OrganismId uint64
}

