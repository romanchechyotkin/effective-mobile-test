package dbo

type User struct {
	ID          string `db:"id"`
	LastName    string `db:"last_name"`
	FirstName   string `db:"first_name"`
	SecondName  string `db:"second_name"`
	Age         int    `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}

func NewUser(lastName, firstName, secondName string, age int, gender, nation string) *User {
	return &User{
		LastName:    lastName,
		FirstName:   firstName,
		SecondName:  secondName,
		Age:         age,
		Gender:      gender,
		Nationality: nation,
	}
}
