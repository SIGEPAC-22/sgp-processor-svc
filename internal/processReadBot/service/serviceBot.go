package service

import (
	"context"
	"github.com/go-kit/log"
	"sgp-processor-svc/internal/processReadBot"
	"sgp-processor-svc/kit/constants"
	"sgp-processor-svc/kit/saveHistoryPaientInfo"
)

type serviceDataHistorical struct {
	repo processReadBot.RepositoryDataHistorical
	log  log.Logger
}

func NewServiceDataHistorical(repo processReadBot.RepositoryDataHistorical, log log.Logger) *serviceDataHistorical {
	return &serviceDataHistorical{repo: repo, log: log}
}

func (s *serviceDataHistorical) ServiceProcessGetDataHistorical(ctx context.Context, request []saveHistoryPaientInfo.GetDataHistoricalResponse) error {

	for _, data := range request {

		errRepo := s.repo.InsertProcessGetDataHistorical(ctx, data.IdPatient, data.IdPatientFile, data.FirstName, data.SecondName, data.FirstLastName, data.SecondLastName, data.AdmissionDate, data.HighDate, data.LowDate)
		if errRepo != nil {
			s.log.Log("Repo insert failed", constants.UUID, ctx.Value(constants.UUID))
			continue
		}
	}
	return nil
}
