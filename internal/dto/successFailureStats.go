package dto

import "sync"

type SuccessFailureStats struct {
	Successes uint32 `json:"successes"`
	Failures  uint32 `json:"failures"`

	mu sync.Mutex
}

func (s *SuccessFailureStats) AddSuccesses(count uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Successes += count
}

func (s *SuccessFailureStats) AddFailures(count uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Failures += count
}
