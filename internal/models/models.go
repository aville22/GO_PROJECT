package models

type ProfileForm struct {
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	Age      float64 `json:"age"`
	Gender   string  `json:"gender"`
	Activity float64 `json:"activity"`
	Goal     string  `json:"goal"`
}
