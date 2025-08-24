package models

type OrderItem struct {
	ID            string `json:"id"`
	Quantity      int    `json:"quantity" validate:"required, eq=S|eq=M|eq=L|eq=XL"`
	Unit_price    int    `json:"unit_price" validate:"required"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Food_id       string `json:"food_id" validate:"required"`
	Order_item_id string `json:"order_item_id"`
	Order_id      string `json:"order_id" validate:"required"`
}
