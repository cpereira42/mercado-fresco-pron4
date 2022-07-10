package seller

type Seller struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type SellerRequestCreate struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  int    `json:"locality_id" binding:"required"`
}

type SellerRequestUpdate struct {
	Cid         string `json:"cid" binding:"omitempty,required"`
	CompanyName string `json:"company_name" binding:"omitempty,required"`
	Address     string `json:"address" binding:"omitempty,required"`
	Telephone   string `json:"telephone" binding:"omitempty,required"`
	LocalityId  int    `json:"locality_id" binding:"omitempty,required"`
}
