package entities

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Sales struct {
	ID          int    `json:"salesid"`
	Productname string `json:"productname"`
	Nrofsales   int    `json:"nrofsales"`
}
