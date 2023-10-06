package bq

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/rs/zerolog"
)

type visitRepository struct {
	db     *bigquery.Client
	table  *bigquery.Table
	logger *zerolog.Logger
}

func NewVisitRepository(db *bigquery.Client, datasetID, tableID string, logger *zerolog.Logger) *visitRepository {
	dataset := db.Dataset(datasetID)
	table := dataset.Table(tableID)
	return &visitRepository{
		db:     db,
		table:  table,
		logger: logger,
	}
}

func (vr visitRepository) TableExists(ctx context.Context) error {
	err := TableExists(ctx, vr.table)
	if err != nil {
		return err
	}
	return nil
}

func (vr visitRepository) CreateTable(ctx context.Context, fieldPartition string, fieldClustering []string, schema any) error {
	err := CreateTable(ctx, vr.table, &fieldPartition, fieldClustering, schema)
	if err != nil {
		return fmt.Errorf("createTable: %w", err)
	}
	return nil
}

func (vr visitRepository) DeleteByDateColumn(ctx context.Context, dateFrom, dateTo string) error {
	vr.logger.Trace().Msg("deleteByDateColumn")
	dateColumn := "date"
	err := DeleteByDateColumn(ctx, vr.db, vr.table, dateColumn, dateFrom, dateTo)
	if err != nil {
		return fmt.Errorf("DeleteByDateColumn: %w", err)
	}
	return nil
}

func (vr visitRepository) SendFromCS(ctx context.Context, bucket, object string) error {
	vr.logger.Info().Msg("Load Visits to BQ")
	err := SendFromCS(ctx, vr.table, bucket, object, VisitSchema{})
	if err != nil {
		return fmt.Errorf("SendFromCS: %w", err)
	}
	return nil
}
