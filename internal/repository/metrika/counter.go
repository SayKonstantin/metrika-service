package metrika

import (
	"context"
	"fmt"
	sdk "github.com/mg-realcom/metrika-sdk"
	"github.com/rs/zerolog"
	"strings"
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

func (l *CounterRepository) GetCounters(ctx context.Context) (string, error) {
	counters, err := l.client.GetCounters(ctx)
	if err != nil {
		return "", err
	}
	var res = []string{"counterID | counterName"}
	for _, c := range counters {
		counter := fmt.Sprintf("%v - %v", c.Id, c.Name)
		res = append(res, counter)
	}
	return strings.Join(res, "\n"), nil
}
