package manager

type (
	User struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	}

	Service interface {
		PostUser(user User) error
	}
)
