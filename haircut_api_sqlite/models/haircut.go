package models

type Haircut struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ID          int     `json:"id"`
	Price       float64 `json:"price"`
}
