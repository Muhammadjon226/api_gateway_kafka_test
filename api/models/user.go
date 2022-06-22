package models

//CreateUserModel ...
type CreateUserModel struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//User ..
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//ListUserResponse ...
type ListUserResponse struct {
	Users []*User `json:"users"`
	Count int64   `json:"count"`
}
