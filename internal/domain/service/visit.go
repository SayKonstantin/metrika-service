package service

import (
	"context"
	"fmt"
	"github.com/SayKonstantin/metrika-service/internal/repository/bq"
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
	CreateTable(ctx context.Context, fieldPartition string, fieldClustering []string, schema any) (err error)
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

func (vs *VisitService) GetVisits(ctx context.Context) error {
	vs.logger.Trace().Msg("VisitsToBQ")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, vs.dates.DateFrom)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	endDate, err := time.Parse(layout, vs.dates.DateTo)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	unix1 := startDate.Unix()
	unix2 := endDate.Unix()
	err = vs.bq.TableExists(ctx)
	if err != nil {
		err = vs.bq.CreateTable(ctx, "date", nil, &bq.VisitSchema{})
		if err != nil {
			return fmt.Errorf("can`t create table: %w", err)
		}
		vs.logger.Info().Msg("Visits Table created")
	}
	for i := unix1; i <= unix2; i += 86400 {
		t1 := time.Unix(i, 0)
		dt := i + 86400
		if dt > unix2 {
			dt = unix2
		}
		actualDate := fmt.Sprintf("%d-%02d-%02d", t1.Year(), t1.Month(), t1.Day())
		files, err := vs.metrika.PushLog(ctx, actualDate, actualDate)
		if err != nil {
			return err
		}

		err = vs.bq.DeleteByDateColumn(ctx, vs.dates.DateFrom, vs.dates.DateTo)
		if err != nil {
			return err
		}

		err = vs.PushVisitsToBQ(ctx, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func (vs VisitService) PushVisitsToBQ(ctx context.Context, files []string) error {
	for _, file := range files {
		err := vs.cs.SendFile(ctx, file)
		if err != nil {
			return fmt.Errorf("can`t send file to CS: %w", err)
		}
		bucket, err := vs.cs.GetBucket(ctx)
		if err != nil {
			return err
		}
		err = vs.bq.SendFromCS(ctx, bucket, file)
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
