package models

//CreateUserModel ...
type CreateUserModel struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//UserModel ..
type UserModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}