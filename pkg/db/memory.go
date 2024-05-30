package db

import (
	"encoding/json"
	"fmt"

	"github.com/vengeancegod/tgbot/pkg/util"
)

type MemoryStore struct {
	namespace string
	memory    map[string]string
}

func NewMemoryStore(namespace string) *MemoryStore {
	return &MemoryStore{
		namespace: namespace,
		memory:    make(map[string]string),
	}
}

func (s *MemoryStore) Load(key, value interface{}) error {
	defer util.Timer("Хранилище загружено")()
	data, ok := s.memory[cast(key)]
	if !ok {
		return fmt.Errorf("Ключ %s не найден", key)
	}
	return json.Unmarshal([]byte(data), value)
}

func (s *MemoryStore) Save(key, value interface{}) error {
	defer util.Timer("Хранилище сохранено")()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	s.memory[cast(key)] = string(data)
	return nil
}

func (s *MemoryStore) Delete(key interface{}) error {
	defer util.Timer("Хранилище удалено")()
	delete(s.memory, cast(key))
	return nil
}
