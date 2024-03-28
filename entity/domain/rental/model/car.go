package model

// Car is an entity for car.
type Car struct {
	ID       string   `json:"id"`
	Location Location `json:"location,omitempty"`
}
