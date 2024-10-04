package employer

type CreateEmployerReq struct {
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Address   string    `json:"address"`
    Country   string    `json:"country"`
    Website   string    `json:"website"`
}

type EmployerRes struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Country string `json:"country"`
	Website string `json:"website"`
}