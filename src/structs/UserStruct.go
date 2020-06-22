package structs

type User struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Company *Organism //`json:"company",one-to-many:"Organism"`
}
