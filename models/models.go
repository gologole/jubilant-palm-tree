package models

type Order struct {
	ID          int64  `json:"id"`
	OverPrice   int64  `json:"over_price"`
	Description string `json:"description"`
}

// то что возвращаю на фронтенд
type Document struct {
	Cost      int64     `json:"cost"`
	OverPrice int64     `json:"over_price"`
	Courier   Courier   `json:"courier"`
	Couriers  []Courier `json:"couriers"`
}

type Courier struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Dist   int64  `json:"dist"`
	Status string `json:"status"`
}
