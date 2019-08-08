package utils

import "sync"

type StringSet struct {
	m map[string]bool
	sync.RWMutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		m: map[string]bool{},
	}
}
func (s *StringSet) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *StringSet) Remove(item string) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *StringSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *StringSet) Len() int {
	return len(s.List())
}

func (s *StringSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *StringSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *StringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

type IntSet struct {
	m map[int]bool
	sync.RWMutex
}

func NewIntSet() *IntSet {
	return &IntSet{
		m: map[int]bool{},
	}
}
func (s *IntSet) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}
func (s *IntSet) Remove(item int) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}
func (s *IntSet) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *IntSet) Len() int {
	return len(s.List())
}

func (s *IntSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *IntSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *IntSet) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
