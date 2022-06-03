package seller

type Seller struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Adress      string `json:"address"`
	Telephone   string `json:"telephone"`
}

type Repository interface {
	GetAll() ([]Seller, error)
	GetId(id int) (Seller, error)
	Create(id, cid int, company, adress, telephone string) (Seller, error)
	LastID() (int, error)
	//Update(id int, prod Product)
	//Delete(id int) error
}
