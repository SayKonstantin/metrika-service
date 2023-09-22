package policy

import (
	"context"
	"github.com/SayKonstantin/metrika-service/internal/domain/service"
)

type VisitPolicy struct {
	Service service.VisitService
}

func NewVisitPolicy(srv service.VisitService) *VisitPolicy {
	return &VisitPolicy{Service: srv}

}

func (mp VisitPolicy) PushVisits(ctx context.Context) error {
	err := mp.Service.GetVisits(ctx)
	if err != nil {
		return err
	}
	return nil
}
