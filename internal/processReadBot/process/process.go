package process

import (
	"context"
	"github.com/go-kit/log"
	"sgp-processor-svc/internal/processReadBot"
	"sgp-processor-svc/kit/constants"
	"sgp-processor-svc/kit/saveHistoryPaientInfo"
)

type processReader struct {
	log        log.Logger
	repository saveHistoryPaientInfo.Repository
	service    processReadBot.ServiceDataHistorical
}

func NewProcessReader(log log.Logger, repository saveHistoryPaientInfo.Repository, service processReadBot.ServiceDataHistorical) *processReader {
	return &processReader{log: log, repository: repository, service: service}
}

func (p processReader) ProcessReader(ctx context.Context) {
	respRepo, errRepo := p.repository.GetDataPersonal(ctx)
	if errRepo != nil {
		p.log.Log("Repo Failed", constants.UUID, ctx.Value(constants.UUID))
	}

	errService := p.service.ServiceProcessGetDataHistorical(ctx, respRepo)
	if errService != nil {

	}
}
