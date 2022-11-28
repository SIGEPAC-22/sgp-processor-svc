package process

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-processor-svc/internal/processReadBot"
	"sgp-processor-svc/kit/saveHistoryPaientInfo"
	"time"
)

type ProcessReaderBot struct {
	ctx        context.Context
	logger     log.Logger
	repository saveHistoryPaientInfo.Repository
	service    processReadBot.ServiceDataHistorical
}

func NewProcessReaderBot(ctx context.Context, logger log.Logger, repository saveHistoryPaientInfo.Repository, service processReadBot.ServiceDataHistorical) *ProcessReaderBot {
	return &ProcessReaderBot{ctx: ctx, logger: logger, repository: repository, service: service}
}

func (p *ProcessReaderBot) ProcessReaderInitializer() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	p.logger.Log("Process Reader bot", "mensaje", "Iniciando Robot")
	loc, _ := time.LoadLocation(config.GetString("queue-bot-process.dateLocalFormat"))
	s := gocron.NewScheduler(loc)

	process := NewProcessReader(p.logger, p.repository, p.service)
	s.Every(config.GetString("queue-bot-process.time-run")).Do(process.ProcessReader, context.Background())
	s.StartAsync()
}
