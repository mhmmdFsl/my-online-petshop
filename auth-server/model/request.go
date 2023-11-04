package model

type SignUpRq struct {
	Name      string `json:"name"`
	Principal string `json:"principal"`
	Password  string `json:"password"`
}

type LoginRq struct {
	Principal string `json:"principal"`
	Password  string `json:"password"`
}
