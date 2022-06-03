package products

type Product struct {
	Id                             int     `json:"id"`
	Product_code                   string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Length                         float64 `json:"length"`
	Height                         float64 `json:"height"`
	NetHeight                      float64 `json:"net_height"`
	ExpirationRate                 string  `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductType_Id                 int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type Repository interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	Delete(id int) error
	LastID() (int, error)
	Store(p Product) (Product, error)
	Update(id int, prod Product) (Product, error)
	UpdatePatch(id int, prod Product) (Product, error)
}
