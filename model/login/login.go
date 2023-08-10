package model_login

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
