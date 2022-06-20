package seller

type Seller struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Adress      string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerRequestCreate struct {
	Cid         int    `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Adress      string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
}

type SellerRequestUpdate struct {
	Cid         int    `json:"cid" binding:"omitempty,required"`
	CompanyName string `json:"company_name" binding:"omitempty,required"`
	Adress      string `json:"address" binding:"omitempty,required"`
	Telephone   string `json:"telephone" binding:"omitempty,required"`
}
