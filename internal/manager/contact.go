package manager

type (
	User struct {
		Sity   string `json:"sity"`
		Name   string `json:"name"`
		Number string `json:"number"`
	}

	Service interface {
		PostUser(user User) error
	}
)
