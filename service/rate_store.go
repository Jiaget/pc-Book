package service

import "sync"

// RateStore is an interface to store laotop ratings
type RateStore interface {
	// Add adds a new laptop score to the store and return the rate
	Add(laptopID string, score float64) (*Rate, error)
}

// Rate contains the rating information of a laptop
type Rate struct {
	Count uint32
	Sum   float64
}

// InMemoryRateStore stores laptop rate in memory
type InMemoryRateStore struct {
	mutex sync.RWMutex
	rate  map[string]*Rate
}

// NewInMemoryRateStore creates a new InMemoryRateStore
func NewInMemoryRateStore() *InMemoryRateStore {
	return &InMemoryRateStore{
		rate: make(map[string]*Rate),
	}
}

func (store *InMemoryRateStore) Add(laptopID string, score float64) (*Rate, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rate := store.rate[laptopID]
	if rate == nil {
		rate = &Rate{
			Count: 1,
			Sum:   score,
		}
	} else {
		rate.Count++
		rate.Sum += score
	}
	store.rate[laptopID] = rate
	return rate, nil
}
