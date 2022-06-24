package products

type Product struct {
	Id                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Length                         float64 `json:"length"`
	Height                         float64 `json:"height"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type RequestProductsCreate struct {
	ProductCode                    string  `json:"product_code" binding:"required"`
	Description                    string  `json:"description" binding:"required"`
	Width                          float64 `json:"width" binding:"required"`
	Length                         float64 `json:"length" binding:"required"`
	Height                         float64 `json:"height" binding:"required"`
	NetWeight                      float64 `json:"net_weight" binding:"required"`
	ExpirationRate                 float64 `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate                   float64 `json:"freezing_rate" binding:"required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"required"`
	SellerId                       int     `json:"seller_id" binding:"omitempty,required"`
}

type RequestProductsUpdate struct {
	ProductCode                    string  `json:"product_code" binding:"omitempty,required"`
	Description                    string  `json:"description" binding:"omitempty,required"`
	Width                          float64 `json:"width" binding:"omitempty,required"`
	Length                         float64 `json:"length" binding:"omitempty,required"`
	Height                         float64 `json:"height" binding:"omitempty,required"`
	NetWeight                      float64 `json:"net_weight" binding:"omitempty,required"`
	ExpirationRate                 float64 `json:"expiration_rate" binding:"omitempty,required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" binding:"omitempty,required"`
	FreezingRate                   float64 `json:"freezing_rate" binding:"omitempty,required"`
	ProductTypeId                  int     `json:"product_type_id" binding:"omitempty,required"`
	SellerId                       int     `json:"seller_id" binding:"omitempty,required"`
}
