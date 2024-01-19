package httpsrv

type UserRequest struct {
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name,omitempty"`
}

type UserResponse struct {
	ID          string `json:"id"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name,omitempty"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type UpdateUserDto struct {
	LastName    string `json:"last_name,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	SecondName  string `json:"second_name,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}
