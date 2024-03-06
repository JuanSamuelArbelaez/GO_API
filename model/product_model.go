package model

type Product struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    UnitValue float32 `json:"value"`
    Units     int     `json:"units"`
}