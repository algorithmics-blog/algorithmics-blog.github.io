package unique_value_structure

import (
	"math/rand"
)

type Store struct {
	indexesMap map[int]int
	values     []int
}

func New() *Store {
	return &Store{
		indexesMap: make(map[int]int),
		values:     make([]int, 0),
	}
}

func (s *Store) Insert(value int) {
	_, exists := s.indexesMap[value]
	if exists {
		return
	}

	s.values = append(s.values, value)
	s.indexesMap[value] = len(s.values) - 1
}

func (s *Store) Remove(value int) {
	index, exists := s.indexesMap[value]
	if !exists {
		return
	}

	last := s.values[len(s.values)-1]
	s.values[index] = last
	s.indexesMap[last] = index

	s.values = s.values[:len(s.values)-1]
	delete(s.indexesMap, value)
}

func (s *Store) GetRandom() int {
	index := rand.Intn(len(s.values))
	return s.values[index]
}

// Method for tests
func (s *Store) getValues() []int {
	return s.values
}
