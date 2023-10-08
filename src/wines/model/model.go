package model

type Wine struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Grape       string   `json:"grape"`
	Vintage     int      `json:"vintage"`
	Region      string   `json:"region"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	Sweetness   string   `json:"sweetness"`
	Acidity     string   `json:"acidity"`
	Tannins     string   `json:"tannins"`
	Alcohol     float64  `json:"alcohol"`
	Finish      string   `json:"finish"`
	Pairing     []string `json:"pairing"`
}
