package model

type Person struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Country   string `json:"country"`
	State     string `json:"state"`
}

type PersonRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Country   string `json:"country"`
}

func MapRequest(p *Person, r *PersonRequest) {
	p.Name = r.Name
	p.LastName = r.LastName
	p.Email = r.Email
	p.Telephone = r.Telephone
	p.Country = r.Country
}
