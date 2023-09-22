package bq

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type logRepository struct {
	db     *bigquery.Client
	table  *bigquery.Table
	logger *zerolog.Logger
}

func NewMetrikaRepository(db *bigquery.Client, datasetID, tableID string, logger *zerolog.Logger) *logRepository {
	dataset := db.Dataset(datasetID)
	table := dataset.Table(tableID)
	return &logRepository{
		db:     db,
		table:  table,
		logger: logger,
	}
}

func (mr logRepository) TableExists(ctx context.Context) error {
	err := TableExists(ctx, mr.table)
	if err != nil {
		return err
	}
	return nil
}

func (mr logRepository) CreateTable(ctx context.Context, schema any) error {
	fieldPartition := "date"
	//fieldClasteting := []string{"ad_id", "campaign_id"}
	//fieldClasteting := []string{}
	err := CreateTable(ctx, mr.table, &fieldPartition, nil, schema)
	if err != nil {
		return fmt.Errorf("createTable: %w", err)
	}
	return nil
}

func (mr logRepository) DeleteByDateColumn(ctx context.Context, dateFrom, dateTo string) error {
	mr.logger.Trace().Msg("deleteByDateColumn")
	dateColumn := "date"
	err := DeleteByDateColumn(ctx, mr.db, mr.table, dateColumn, dateFrom, dateTo)
	if err != nil {
		return fmt.Errorf("DeleteByDateColumn: %w", err)
	}
	return nil
}

func (mr logRepository) SendFromCS(ctx context.Context, bucket, object string) error {
	err := SendFromCS(ctx, mr.table, bucket, object)
	if err != nil {
		return fmt.Errorf("SendFromCS: %w", err)
	}
	return nil
}
