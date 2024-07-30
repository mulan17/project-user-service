package profile

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Address  Address
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	House   string `json:"house"`
}
