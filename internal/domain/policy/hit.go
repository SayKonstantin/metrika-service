package policy

import (
	"context"
	"github.com/SayKonstantin/metrika-service/internal/domain/service"
)

type HitPolicy struct {
	Service service.HitService
}

func NewHitPolicy(srv service.HitService) *HitPolicy {
	return &HitPolicy{Service: srv}

}

func (mp HitPolicy) PushHitsToBQ(ctx context.Context) error {
	err := mp.Service.GetHits(ctx)
	if err != nil {
		return err
	}
	return nil
}
