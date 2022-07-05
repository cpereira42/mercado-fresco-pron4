package locality

type Locality struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type LocalityRequestCreate struct {
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
}

type LocalityRequestUpdate struct {
	LocalityName string `json:"locality_name" binding:"omitempty,required"`
	ProvinceName string `json:"province_name" binding:"omitempty,required"`
	CountryName  string `json:"country_name" binding:"omitempty,required"`
}

type LocalityReport struct {
	LocalityId   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}
