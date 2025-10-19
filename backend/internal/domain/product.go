package domain

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"` // dalam satuan terkecil (misal: rupiah)
}
