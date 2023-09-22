package cs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

type logRepository struct {
	bucket    *storage.BucketHandle
	client    *storage.Client
	logger    *zerolog.Logger
	projectId string
}

func NewMetrikaRepository(bucketName, projectId string, client *storage.Client, logger *zerolog.Logger) *logRepository {
	bucket := client.Bucket(bucketName)
	return &logRepository{
		bucket:    bucket,
		client:    client,
		logger:    logger,
		projectId: projectId,
	}
}

func (mr logRepository) SendFile(ctx context.Context, filename string) error {
	mr.logger.Info().Msgf("SendFile: %v", filename)
	_, err := mr.bucket.Attrs(ctx)

	if err != nil {
		err = mr.bucket.Create(ctx, mr.projectId, &storage.BucketAttrs{RetentionPolicy: &storage.RetentionPolicy{RetentionPeriod: 5 * 24 * time.Hour}})
		if err != nil {
			return fmt.Errorf("bucket creation error: %w", err)
		}
	}
	err = SendFile(ctx, mr.bucket, filename)
	if err != nil {
		return fmt.Errorf("SendFile error: %w", err)
	}

	return nil
}
func (mr logRepository) GetBucket(ctx context.Context) (string, error) {
	bucketName, err := mr.bucket.Attrs(ctx)
	if err != nil {
		return "", err
	}
	return bucketName.Name, nil

}
