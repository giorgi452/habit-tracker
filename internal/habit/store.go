package habit

import (
	"fmt"
	"sync"
)

type Store struct {
	mu     sync.RWMutex
	habits []*Habit
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Add(
	name string,
	freq Frequency,
	durationStr string,
	timeStrings []string,
) (*Habit, error) {
	h, err := newHabit(name, freq, durationStr, timeStrings)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	h.ID = len(s.habits) + 1
	s.habits = append(s.habits, h)
	return h, nil
}

func (s *Store) Range(fn func(*Habit)) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, h := range s.habits {
		fn(h)
	}
}

func (s *Store) List() []*Habit {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.habits
}

func (s *Store) Get(id int) (*Habit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, h := range s.habits {
		if h.ID == id {
			return h, nil
		}
	}
	return nil, fmt.Errorf("Habit with ID %d not found", id)
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, h := range s.habits {
		if h.ID == id {
			s.habits = append(s.habits[:i], s.habits[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Habit with ID %d not found", id)
}

func (s *Store) Update(id int, fn func(*Habit) error) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, h := range s.habits {
		if h.ID == id {
			return fn(h)
		}
	}
	return fmt.Errorf("Habit with ID %d not found", id)
}
