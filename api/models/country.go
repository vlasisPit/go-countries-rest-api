package models

type Country struct {
	Name       string     `json:"name"`
	Alpha2Code string     `json:"alpha2Code"`
	Capital    string     `json:"capital"`
	Currencies []Currency `json:"currencies"`
}