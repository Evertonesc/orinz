package member

type Member struct {
	Name      string
	BirthDate string
	Address   string //Make a struct
	Email     string //Make a struct
}

func New(name, birthDate, address, email string) Member {
	return Member{
		Name:      name,
		BirthDate: birthDate,
		Address:   address,
		Email:     email,
	}
}
