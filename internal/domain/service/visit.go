package service

import (
	"context"
	"fmt"
	"github.com/SayKonstantin/metrika-service/internal/repository/metrika"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type VisitRepository interface {
	PushLog(ctx context.Context, dateFrom, dateTo string) ([]string, error)
}

type VisitCsRepository interface {
	SendFile(ctx context.Context, filename string) (err error)
	GetBucket(ctx context.Context) (string, error)
}

type VisitBqRepository interface {
	SendFromCS(ctx context.Context, bucket string, object string) (err error)
	DeleteByDateColumn(ctx context.Context, dateFrom, dateTo string) (err error)
	CreateTable(ctx context.Context, schema any) (err error)
	TableExists(ctx context.Context) (err error)
}

type VisitService struct {
	metrika VisitRepository
	cs      VisitCsRepository
	bq      VisitBqRepository
	dates   metrika.DateRange
	logger  *zerolog.Logger
}

func NewVisitService(metrika VisitRepository, cs VisitCsRepository, bq VisitBqRepository, dates metrika.DateRange, logger *zerolog.Logger) *VisitService {
	return &VisitService{
		metrika: metrika,
		cs:      cs,
		bq:      bq,
		dates:   dates,
		logger:  logger,
	}
}

func (ms *VisitService) GetVisits(ctx context.Context) error {
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, ms.dates.DateFrom)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	endDate, err := time.Parse(layout, ms.dates.DateTo)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	unix1 := startDate.Unix()
	unix2 := endDate.Unix()
	for i := unix1; i <= unix2; i += 86400 {
		t1 := time.Unix(i, 0)
		dt := i + 86400
		if dt > unix2 {
			dt = unix2
		}
		actualDate := fmt.Sprintf("%d-%02d-%02d", t1.Year(), t1.Month(), t1.Day())
		files, err := ms.metrika.PushLog(ctx, actualDate, actualDate)
		if err != nil {
			return err
		}
		err = ms.PushVisitsToBQ(ctx, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms VisitService) PushVisitsToBQ(ctx context.Context, files []string) error {
	for _, file := range files {
		err := ms.cs.SendFile(ctx, file)
		if err != nil {
			return fmt.Errorf("can`t send file to CS: %w", err)
		}
		bucket, err := ms.cs.GetBucket(ctx)
		if err != nil {
			return err
		}
		err = ms.bq.SendFromCS(ctx, bucket, file)
		if err != nil {
			return fmt.Errorf("can`t send file from CS: %w", err)
		}
		err = os.Remove(file)
		if err != nil {
			return fmt.Errorf("can`t remove file: %w", err)
		}
	}
	return nil
}
