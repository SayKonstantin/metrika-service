package policy

import (
	"context"
	"github.com/SayKonstantin/metrika-service/internal/domain/service"
)

type CounterPolicy struct {
	Service service.CounterService
}

func NewCounterPolicy(srv service.CounterService) *CounterPolicy {
	return &CounterPolicy{Service: srv}

}

func (mp CounterPolicy) GetCounters(ctx context.Context) (string, error) {
	counters, err := mp.Service.GetCounters(ctx)
	if err != nil {
		return "", err
	}
	return counters, nil
}
