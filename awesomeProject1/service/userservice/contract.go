package userservice

import "awesomeProject1/entity"

type MemoryStorage interface {
	Update(entity.Representation) error
	GetRepByID(int) (entity.Representation, error)
	Last() int
	ListAll() ([]entity.Representation, error)
	Save(entity.Representation) error
	Status(string) (int, int, error)
}

type repository interface {
	Update(MemoryStorage) error
	Save(entity.Representation) error
	Load() ([]entity.Representation, error)
}

type reader interface {
	Create() (entity.Representation, error)
	Edit(MemoryStorage, string) (entity.Representation, error)
	CommandRegion() (string, string)
	GetID(MemoryStorage) (int, error)
}

type writer interface {
	GetRepByID(entity.Representation) error
	GetStatus(int, int) error
	ListByRegion([]entity.Representation) error
}

type Service struct {
	Memory     MemoryStorage
	Repository repository
	Reader     reader
	Writer     writer
}

func New(memory MemoryStorage, repository repository, reader reader, writer writer) Service {
	srv := Service{
		Memory:     memory,
		Repository: repository,
		Reader:     reader,
		Writer:     writer,
	}

	return srv
}
