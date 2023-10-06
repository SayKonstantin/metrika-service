package cs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

type hitRepository struct {
	bucket    *storage.BucketHandle
	client    *storage.Client
	logger    *zerolog.Logger
	projectId string
}

func NewHitRepository(bucketName, projectId string, client *storage.Client, logger *zerolog.Logger) *hitRepository {
	bucket := client.Bucket(bucketName)
	return &hitRepository{
		bucket:    bucket,
		client:    client,
		logger:    logger,
		projectId: projectId,
	}
}

func (hr hitRepository) SendFile(ctx context.Context, filename string) error {
	hr.logger.Info().Msgf("SendFile to CS: %v", filename)
	_, err := hr.bucket.Attrs(ctx)

	if err != nil {
		err = hr.bucket.Create(ctx, hr.projectId, &storage.BucketAttrs{RetentionPolicy: &storage.RetentionPolicy{RetentionPeriod: 5 * 24 * time.Hour}})
		if err != nil {
			return fmt.Errorf("bucket creation error: %w", err)
		}
	}
	err = SendFile(ctx, hr.bucket, filename)
	if err != nil {
		return fmt.Errorf("SendFile error: %w", err)
	}

	return nil
}
func (hr hitRepository) GetBucket(ctx context.Context) (string, error) {
	bucketName, err := hr.bucket.Attrs(ctx)
	if err != nil {
		return "", err
	}
	return bucketName.Name, nil

}
