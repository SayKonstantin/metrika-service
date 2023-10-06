package cs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

type visitRepository struct {
	bucket    *storage.BucketHandle
	client    *storage.Client
	logger    *zerolog.Logger
	projectId string
}

func NewVisitRepository(bucketName, projectId string, client *storage.Client, logger *zerolog.Logger) *visitRepository {
	bucket := client.Bucket(bucketName)
	return &visitRepository{
		bucket:    bucket,
		client:    client,
		logger:    logger,
		projectId: projectId,
	}
}

func (vr visitRepository) SendFile(ctx context.Context, filename string) error {
	vr.logger.Info().Msgf("SendFile to CS: %v", filename)
	_, err := vr.bucket.Attrs(ctx)

	if err != nil {
		err = vr.bucket.Create(ctx, vr.projectId, &storage.BucketAttrs{RetentionPolicy: &storage.RetentionPolicy{RetentionPeriod: 5 * 24 * time.Hour}})
		if err != nil {
			return fmt.Errorf("bucket creation error: %w", err)
		}
	}
	err = SendFile(ctx, vr.bucket, filename)
	if err != nil {
		return fmt.Errorf("SendFile error: %w", err)
	}

	return nil
}
func (vr visitRepository) GetBucket(ctx context.Context) (string, error) {
	bucketName, err := vr.bucket.Attrs(ctx)
	if err != nil {
		return "", err
	}
	return bucketName.Name, nil

}
