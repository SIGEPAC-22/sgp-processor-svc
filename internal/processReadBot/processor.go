package processReadBot

import (
	"context"
	"sgp-processor-svc/kit/saveHistoryPaientInfo"
)

type IProcessReaderBot interface {
	ProcessReaderInitializer()
}

type ReaderProccesor interface {
	ProcessReader(ctx context.Context)
}

type ServiceDataHistorical interface {
	ServiceProcessGetDataHistorical(ctx context.Context, request []saveHistoryPaientInfo.GetDataHistoricalResponse) error
}

type RepositoryDataHistorical interface {
	InsertProcessGetDataHistorical(ctx context.Context, idPatient int, idPatientFile int, firstName string, secondName string, firstLastName string, secondLastName string, admissionDate string, highDate string, lowDate string) error
}

type GetDataHistoricalResponse struct {
	IdPatient      int    `json:"id"`
	IdPatientFile  int    `json:"idPatientFile"`
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	FirstLastName  string `json:"firstLastName"`
	SecondLastName string `json:"secondLastName"`
	AdmissionDate  string `json:"admissionDate"`
	HighDate       string `json:"highDate"`
	LowDate        string `json:"lowDate"`
}
