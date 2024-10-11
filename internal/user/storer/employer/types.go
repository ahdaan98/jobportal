package employer

type EmployerRes struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Phone   string `db:"phone"`
	Address string `db:"address"`
	Country string `db:"country"`
	Website string `db:"website"`
}

type CreateEmployerReq struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
	Address  string `db:"address"`
	Country  string `db:"country"`
	Website  string `db:"website"`
}

type GetEPI struct {
	Id    int64  `db:"id"`
	Email string `db:"email"`
	Pass  string `db:"password"`
}
