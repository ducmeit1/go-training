package cmd

import (
	"errors"
	"go-training/db"
	"go-training/entity"
	"go-training/internal"
)

type Store struct {
	bootstraped bool
	DataStore   *db.DataStore
}

func NewStore(rows int) (*Store, error) {
	ds, err := db.NewDataStore(rows)
	if err != nil {
		return nil, err
	}
	return &Store{
		DataStore: ds,
	}, nil
}

func (s *Store) Free() {
	s.DataStore = nil
}

func (s *Store) Bootstrap(rows int) error {
	if s.bootstraped {
		return errors.New("data is bootstraped")
	}
	for i := 0; i < rows; i++ {
		key := internal.UUID()
		// No need to handle error because this is bootstrap
		s.DataStore.Add(key, &entity.People{
			Id:      key,
			Name:    key,
			Age:     internal.RandomInt(23, 50),
			Company: "FPT Software",
			Address: "17 Duy Tan",
		})
	}
	return nil
}

func (s *Store) AddPeople(p *entity.People) error {
	if p == nil {
		return errors.New("people must not be empty")
	}
	if p.Name == "" {
		return errors.New("name must not be empty")
	}
	if p.Age < 23 {
		return errors.New("age must > 23")
	}

	return s.DataStore.Add(p.Id, p)
}

func (s *Store) SearchPeople(id string) (*entity.People, error) {
	return s.DataStore.Read(id)
}

func (s *Store) DeleteById(id string) error {
	return s.DataStore.Delete(id)
}

func (s *Store) UpdateById(id string, p *entity.People) error {
	return s.DataStore.Update(id, p)
}

func (s *Store) PrintAsOrder(order internal.Order) {
	s.DataStore.PrintAsOrder(order)
}
