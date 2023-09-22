package v1

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	mp "github.com/SayKonstantin/metrika-service/internal/domain/policy"
	service "github.com/SayKonstantin/metrika-service/internal/domain/service"
	"github.com/SayKonstantin/metrika-service/internal/repository/bq"
	"github.com/SayKonstantin/metrika-service/internal/repository/cs"
	"github.com/SayKonstantin/metrika-service/internal/repository/metrika"
	"github.com/SayKonstantin/metrika-service/pb"
	sdk "github.com/mg-realcom/metrika-sdk"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

func (s Server) PushVisitsToBQ(ctx context.Context, req *pb.PushLogRequest) (*pb.PushLogResponse, error) {
	dates := metrika.DateRange{DateTo: req.MetrikaConfig.Period.DateTo, DateFrom: req.MetrikaConfig.Period.DateFrom}
	metrikaClient := sdk.Client{
		Tr:        &http.Client{Timeout: 20 * time.Minute},
		Token:     req.MetrikaConfig.YandexToken,
		CounterId: req.MetrikaConfig.CounterId,
	}
	metrikaRepo := metrika.NewVisitRepository(&metrikaClient, s.cfg.AttachmentsDir, metrika.VisitsFields, metrika.VisitsSource, s.logger)
	bqClient, err := bigquery.NewClient(ctx, req.BqConfig.ProjectId, option.WithCredentialsFile(req.BqConfig.ServiceKey))
	if err != nil {
		return nil, err
	}
	bqRepo := bq.NewMetrikaRepository(bqClient, req.BqConfig.DatasetId, req.BqConfig.TableId, s.logger)
	csClient, err := storage.NewClient(ctx, option.WithCredentialsFile(req.BqConfig.ServiceKey))
	if err != nil {
		return nil, err
	}
	csRepo := cs.NewMetrikaRepository(req.CsConfig.BucketName, req.BqConfig.ProjectId, csClient, s.logger)
	srv := service.NewVisitService(metrikaRepo, csRepo, bqRepo, dates, s.logger)
	metrikaPolicy := mp.NewVisitPolicy(*srv)
	err = metrikaPolicy.PushVisits(ctx)
	if err != nil {
		s.notifier.Send(ctx, "Metrika", fmt.Sprintf("PushLog: Error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	s.logger.Info().Msg("PushVisits: Success")

	return &pb.PushLogResponse{Success: true}, nil

}
