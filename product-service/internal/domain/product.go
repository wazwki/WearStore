package domain

type Product struct {
	ID          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Stock       int     `bson:"stock"`
	CreatedAt   string  `bson:"created_at"`
	UpdatedAt   string  `bson:"updated_at"`
}
