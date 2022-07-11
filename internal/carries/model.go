package carries

type Carries struct {
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}
type RequestCarriesCreate struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  int    `json:"locality_id" binding:"required"`
}

type Localities struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	Count        int    `json:"carries_count"`
}
