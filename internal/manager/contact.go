package manager

type (
	User struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		City   string `json:"sity"`
	}

	Service interface {
		PostUser(user User) error
	}
)
