package dbschemas

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json: "name"`
	Price       float32 `json: "price"`
	Description string  `json: "description"`
}
