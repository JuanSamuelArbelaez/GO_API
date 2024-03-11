package model

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	UnitValue float32 `json:"value"`
	Units     int     `json:"units"`
}

type ProductRequest struct {
	Name      string  `json:"name"`
	UnitValue float32 `json:"value"`
	Units     int     `json:"units"`
}

func MapRequest(p *Product, r *ProductRequest) {
	p.Name = r.Name
	p.UnitValue = r.UnitValue
	p.Units = r.Units
}
