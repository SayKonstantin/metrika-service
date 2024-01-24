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

type HitRepository interface {
	PushHits(ctx context.Context, dateFrom, dateTo string) ([]string, error)
}

type HitCsRepository interface {
	SendFile(ctx context.Context, filename string) (err error)
	GetBucket(ctx context.Context) (string, error)
}

type HitBqRepository interface {
	SendFromCS(ctx context.Context, bucket string, object string) (err error)
	DeleteByDateColumn(ctx context.Context, dateFrom, dateTo string) (err error)
	CreateTable(ctx context.Context, fieldPartition string, fieldClustering []string, schema any) (err error)
	TableExists(ctx context.Context) (err error)
}

type HitService struct {
	metrika HitRepository
	cs      HitCsRepository
	bq      HitBqRepository
	dates   metrika.DateRange
	logger  *zerolog.Logger
}

func NewHitService(metrika HitRepository, cs HitCsRepository, bq HitBqRepository, dates metrika.DateRange, logger *zerolog.Logger) *HitService {
	return &HitService{
		metrika: metrika,
		cs:      cs,
		bq:      bq,
		dates:   dates,
		logger:  logger,
	}
}

func (hs *HitService) GetHits(ctx context.Context) error {
	hs.logger.Trace().Msg("PushHitsToBQ")
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, hs.dates.DateFrom)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	endDate, err := time.Parse(layout, hs.dates.DateTo)
	if err != nil {
		return fmt.Errorf("Can`t parse date: %w", err)
	}
	unix1 := startDate.Unix()
	unix2 := endDate.Unix()
	err = hs.bq.TableExists(ctx)
	if err != nil {
		err = hs.bq.CreateTable(ctx, "date", nil, &bq.HitSchema{})
		hs.logger.Info().Msgf("Table created")
		if err != nil {
			return fmt.Errorf("can`t create table: %w", err)
		}
	}

	for i := unix1; i <= unix2; i += 86400 {
		t1 := time.Unix(i, 0)
		dt := i + 86400
		if dt > unix2 {
			dt = unix2
		}
		actualDate := fmt.Sprintf("%d-%02d-%02d", t1.Year(), t1.Month(), t1.Day())
		files, err := hs.metrika.PushHits(ctx, actualDate, actualDate)
		if err != nil {
			return err
		}
		hs.logger.Info().Msgf("Hits from %s have been downloaded", actualDate)

		err = hs.bq.DeleteByDateColumn(ctx, actualDate, actualDate)
		if err != nil {
			return err
		}

		err = hs.PushHitsToBQ(ctx, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func (hs HitService) PushHitsToBQ(ctx context.Context, files []string) error {
	bucket, err := hs.cs.GetBucket(ctx)
	if err != nil {
		return err
	}
	for _, file := range files {
		err := hs.cs.SendFile(ctx, file)
		if err != nil {
			return fmt.Errorf("can`t send file to CS: %w", err)
		}
		err = hs.bq.SendFromCS(ctx, bucket, file)
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
