package service

import (
	"context"
	"github.com/rs/zerolog"
)

type CounterRepository interface {
	GetCounters(ctx context.Context) (string, error)
}

type CounterService struct {
	metrika CounterRepository
	logger  *zerolog.Logger
}

func NewCounterService(metrika CounterRepository, logger *zerolog.Logger) *CounterService {
	return &CounterService{
		metrika: metrika,
		logger:  logger,
	}
}

func (ms *CounterService) GetCounters(ctx context.Context) (string, error) {
	counters, err := ms.metrika.GetCounters(ctx)
	if err != nil {
		return "", err
	}

	return counters, nil
}
