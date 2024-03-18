package station

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/SaimonWoidig/ctc-gas-station/pkg/utils"
	"golang.org/x/sync/semaphore"
)

type Station struct {
	MaxCarsCount         int
	FuelType             string
	semaphore            *semaphore.Weighted
	currentCarsCount     *atomic.Int32
	totalCarsServedCount *atomic.Int32
	minServeDuration     time.Duration
	maxServeDuration     time.Duration
}

func NewStation(size int, minServeDuration, maxServeDuration time.Duration, fuelType string) *Station {
	s := new(Station)
	s.MaxCarsCount = size
	s.FuelType = fuelType
	s.currentCarsCount = &atomic.Int32{}
	s.totalCarsServedCount = &atomic.Int32{}
	s.semaphore = semaphore.NewWeighted(int64(size))
	s.minServeDuration = minServeDuration
	s.maxServeDuration = maxServeDuration
	return s
}

func (s *Station) Serve() (time.Duration, error) {
	if err := s.semaphore.Acquire(context.Background(), 1); err != nil {
		return 0, err
	}
	defer s.semaphore.Release(1)
	s.currentCarsCount.Add(1)
	defer s.currentCarsCount.Add(-1)
	servingTime := utils.RandDurationBetween(s.minServeDuration, s.maxServeDuration)
	time.Sleep(servingTime)
	s.totalCarsServedCount.Add(1)
	return servingTime, nil
}

func (s *Station) IsFull() bool {
	return int(s.currentCarsCount.Load()) >= s.MaxCarsCount
}
