package manager

import wayToDb "github.com/btussupb/vakio/storage"

type service struct {
	wayToDb wayToDb.Storage
}

func (s *service) Exec(userInput User) error {
	return s.wayToDb.Exec(userInput.toStorageModel())
}
