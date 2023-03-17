package manager

import wayToDb "github.com/btussupb/vakio/internal/storage"

func (u User) toStorageModel() wayToDb.User {
	return wayToDb.User{
		Name:   u.Name,
		Number: u.Number,
	}
}
