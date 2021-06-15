package db

import (
	"errors"
	"fmt"
	"go-training/entity"
	"go-training/internal"
	"sort"
)

type DataStore struct {
	store  map[string]*entity.People
	length int
}

func NewDataStore(rows int) (*DataStore, error) {
	if rows <= 0 {
		return nil, errors.New("rows must be > 0")
	}
	return &DataStore{
		store:  make(map[string]*entity.People, rows),
		length: rows,
	}, nil
}

func (ds *DataStore) Add(key string, val *entity.People) error {
	if key == "" {
		return errors.New("key must not be empty")
	}
	if _, ok := ds.store[key]; ok {
		return fmt.Errorf("key: %s is duplicated, the value need a unquie key", key)
	}
	if len(ds.store) == ds.length {
		return fmt.Errorf("the data store is full, only store %d rows", ds.length)
	}

	ds.store[key] = val
	return nil
}

func (ds *DataStore) Update(key string, val *entity.People) error {
	if _, ok := ds.store[key]; !ok {
		return fmt.Errorf("key: %s is not existed to update", key)
	}
	ds.store[key] = val
	return nil
}

func (ds *DataStore) Read(key string) (*entity.People, error) {
	if _, ok := ds.store[key]; !ok {
		return nil, fmt.Errorf("key: %s is not existed to read", key)
	}
	return ds.store[key], nil
}

func (ds *DataStore) Delete(key string) error {
	if _, ok := ds.store[key]; !ok {
		return fmt.Errorf("key: %s is not existed to delete", key)
	}
	delete(ds.store, key)
	return nil
}

func (ds *DataStore) PrintAsOrder(order internal.Order) {
	keys := make([]string, 0)
	for k := range ds.store {
		keys = append(keys, k)
	}
	switch order {
	case internal.Asc:
		sort.Strings(keys)
	case internal.Desc:
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	for _, v := range keys {
		fmt.Println(ds.store[v].ToString())
	}
}
