package manager

import wayToDb "github.com/btussupb/vakio/internal/storage"

func (u User) toStorageModel() wayToDb.User {
	return wayToDb.User{
		City:   u.City,
		Name:   u.Name,
		Number: u.Number,
	}
}
