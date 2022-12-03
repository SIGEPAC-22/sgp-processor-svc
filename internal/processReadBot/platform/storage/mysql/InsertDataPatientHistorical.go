package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"sgp-processor-svc/kit/constants"
)

type GetDataHistorical struct {
	db  *sql.DB
	log log.Logger
}

func NewGetDataHistorical(db *sql.DB, log log.Logger) *GetDataHistorical {
	return &GetDataHistorical{db: db, log: log}
}

func (g *GetDataHistorical) InsertProcessGetDataHistorical(ctx context.Context, idPatient int, idPatientFile int, firstName string, secondName string, firstLastName string, secondLastName string, admissionDate string, highDate string, lowDate string) error {
	sql, err := g.db.ExecContext(ctx, "INSERT INTO his_historical(his_id_patient,his_id_file_patient,his_first_name,his_second_name,his_first_last_name,his_second_last_name,\nhis_admission_date,his_high_date,his_low_date)VALUES(?,?,?,?,?,?,?,?,?);", idPatient, idPatientFile, firstName, secondName, firstLastName, secondLastName, admissionDate, highDate, lowDate)
	g.log.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		g.log.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			g.log.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return err
		}
	}
	return nil
}
