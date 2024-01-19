package bq

import (
	"cloud.google.com/go/bigquery"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/googleapi"
	"net/http"
	"strings"
)

const DeleteByDateQuery = "DELETE `%s.%s` WHERE %s >= '%s' AND %s <= '%s'"

func CreateTable(ctx context.Context, table *bigquery.Table, fieldPartition *string, fieldClustering []string, schemaDTO any) error {
	schema, err := bigquery.InferSchema(schemaDTO)
	if err != nil {
		return fmt.Errorf("bigquery.InferSchema: %w", err)
	}
	metadata := &bigquery.TableMetadata{
		Schema: schema,
	}
	if fieldPartition != nil {
		partition := bigquery.TimePartitioning{
			Type:  bigquery.DayPartitioningType,
			Field: *fieldPartition,
		}
		metadata.TimePartitioning = &partition
	}
	if fieldClustering != nil {
		clustering := bigquery.Clustering{
			Fields: fieldClustering,
		}
		metadata.Clustering = &clustering
	}
	if err := table.Create(ctx, metadata); err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}
	return nil
}

func DeleteByDateColumn(ctx context.Context, bqClient *bigquery.Client, table *bigquery.Table, dateColumn string, dateFrom, dateTo string) error {
	q := bqClient.Query(fmt.Sprintf(DeleteByDateQuery, table.DatasetID, table.TableID, dateColumn, dateFrom, dateColumn, dateTo))
	job, err := q.Run(ctx)
	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

func SendFromCS(ctx context.Context, table *bigquery.Table, bucket, object string, schemaDTO any) error {
	schema, err := bigquery.InferSchema(schemaDTO)
	if err != nil {
		return fmt.Errorf("bigquery.InferSchema: %w", err)
	}
	filePath := strings.Split(object, "/")
	gcsRef := bigquery.NewGCSReference(fmt.Sprintf("gs://%s/%s", bucket, filePath[len(filePath)-1]))
	gcsRef.SourceFormat = bigquery.CSV
	gcsRef.Schema = schema
	gcsRef.CSVOptions.FieldDelimiter = "\t"
	gcsRef.SkipLeadingRows = 1
	loader := table.LoaderFrom(gcsRef)
	loader.WriteDisposition = bigquery.WriteAppend

	loader.CreateDisposition = bigquery.CreateIfNeeded
	job, err := loader.Run(ctx)
	if err != nil {
		return fmt.Errorf("loader error: %w", err)
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	return nil
}

func TableExists(ctx context.Context, table *bigquery.Table) error {
	if _, err := table.Metadata(ctx); err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				return errors.New("dataset or table not found")
			}
		}
	}
	return nil
}
