package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/SayKonstantin/metrika-service/internal/app/server"
	"github.com/SayKonstantin/metrika-service/internal/config"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func main() {
	var fileConfig = flag.String("f", "config.yml", "configuration file")
	var useEnv = flag.Bool("env", true, "use environment variables")

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()
	if !*useEnv {
		logger.Info().Msgf("configuration file: %s", *fileConfig)
	} else {
		logger.Info().Msg("configuration from ENV")
	}
	cfg, err := config.NewServerConfig(*fileConfig, *useEnv)
	appNotify := notify.New()
	telegramService, err := telegram.New(cfg.TG.Token)
	if err != nil {
		logger.Err(err).Msg("Ошибка создания сервиса telegram")
	}
	appNotify.UseServices(telegramService)
	telegramService.AddReceivers(cfg.Chat)
	a := server.NewApp(cfg, &logger, appNotify)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			logger.Error().Msg(fmt.Sprintf("Recovered from panic: %v", r))
		}
	}()
	err = a.Run(ctx)
	if err != nil {
		logger.Err(err).Msg("Ошибка выполнения запроса")
	}

}
