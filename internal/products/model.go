package products

type Product struct {
	Id                             int     `json:"id"`
	Product_code                   int     `json:"ui"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	NetHeight                      float64 `json:"net_height"`
	ExpirationRate                 string  `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductType_Id                 float64 `json:"product_type_id"`
	SellerId                       float64 `json:"seller_id"`
}

type Repository interface {
	GetAll() ([]Product, error)
	GetId(id int) (Product, error)
	//Store(prod Product)
	//Update(id int, prod Product)
	//Delete(id int) error
}
