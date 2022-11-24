package saveHistoryPaientInfo

import "context"

type Repository interface {
	GetDataPersonal(ctx context.Context) ([]GetDataHistoricalResponse, error)
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
