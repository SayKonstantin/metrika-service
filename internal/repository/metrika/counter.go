package metrika

import (
	"context"
	sdk "github.com/mg-realcom/metrika-sdk"
	"github.com/rs/zerolog"
)

type CounterRepository struct {
	client *sdk.Client
	logger *zerolog.Logger
}

func NewCounterRepository(client *sdk.Client, logger *zerolog.Logger) *CounterRepository {
	return &CounterRepository{
		client: client,
		logger: logger,
	}
}

func (l *CounterRepository) GetCounters(ctx context.Context) error {
	counters, err := l.client.GetCounters(ctx)
	if err != nil {
		return err
	}
	for _, c := range counters {
		l.logger.Info().Msgf("counter: %v", c)
	}
	return nil
}
