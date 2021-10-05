package types

type Funk struct {
	Name    string  `json:"name"`
	Service string  `json:"service"`
	SubFunk SubFunk `json:"subfunk"`
}

type SubFunk struct {
	Method string   `json:"method"`
	Geese  []string `json:"geeze"`
}
