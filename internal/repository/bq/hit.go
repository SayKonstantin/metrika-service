package bq

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type hitRepository struct {
	db     *bigquery.Client
	table  *bigquery.Table
	logger *zerolog.Logger
}

func NewHitRepository(db *bigquery.Client, datasetID, tableID string, logger *zerolog.Logger) *hitRepository {
	dataset := db.Dataset(datasetID)
	table := dataset.Table(tableID)
	return &hitRepository{
		db:     db,
		table:  table,
		logger: logger,
	}
}

func (hr hitRepository) TableExists(ctx context.Context) error {
	err := TableExists(ctx, hr.table)
	if err != nil {
		return err
	}
	return nil
}

func (hr hitRepository) CreateTable(ctx context.Context, fieldPartition string, fieldClustering []string, schema any) error {
	err := CreateTable(ctx, hr.table, &fieldPartition, fieldClustering, schema)
	if err != nil {
		return fmt.Errorf("createTable: %w", err)
	}
	return nil
}

func (hr hitRepository) DeleteByDateColumn(ctx context.Context, dateFrom, dateTo string) error {
	hr.logger.Trace().Msg("deleteByDateColumn")
	dateColumn := "date"
	err := DeleteByDateColumn(ctx, hr.db, hr.table, dateColumn, dateFrom, dateTo)
	if err != nil {
		return fmt.Errorf("DeleteByDateColumn: %w", err)
	}
	return nil
}

func (hr hitRepository) SendFromCS(ctx context.Context, bucket, object string) error {
	hr.logger.Info().Msg("Load Hits to BQ")
	err := SendFromCS(ctx, hr.table, bucket, object, HitSchema{})
	if err != nil {
		return fmt.Errorf("SendFromCS: %w", err)
	}
	return nil
}
