package process

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"time"
)

type ProcessReaderBot struct {
	ctx    context.Context
	logger log.Logger
}

func NewProcessReaderBot(ctx context.Context, logger log.Logger) *ProcessReaderBot {
	return &ProcessReaderBot{ctx: ctx, logger: logger}
}

func (p *ProcessReaderBot) ProcessReaderInitializer() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	p.logger.Log("Process Reader bot", "mensaje", "Iniciando Robot")
	loc, _ := time.LoadLocation(config.GetString("queue-bot-process.dateLocalFormat"))
	s := gocron.NewScheduler(loc)

	process := NewProcessReader(p.logger)
	s.Every(config.GetString("queue-bot-process.time-run")).Do(process.ProcessReader, context.Background())
	s.StartAsync()
}
