package saveHistoryPaientInfo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/log"
	"github.com/go-resty/resty/v2"
	goconfig "github.com/iglin/go-config"
	"sgp-processor-svc/kit/constants"
)

type getDataRepository struct {
	client *resty.Client
	log    log.Logger
}

func NewGetDataRepository(client *resty.Client, log log.Logger) *getDataRepository {
	return &getDataRepository{client: client, log: log}
}

func (g *getDataRepository) GetDataPersonal(ctx context.Context) ([]GetDataHistoricalResponse, error) {

	var response []GetDataHistoricalResponse

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	url := config.GetString("service.sgp-info-svc.getDataHistorical")

	g.log.Log("Get Data Historical Patient", "URL request", url, constants.UUID, ctx.Value(constants.UUID))

	g.log.Log("GetDataHistorical - Repo - Request", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.client.R().
		SetHeader("Content-Type", "application/json").
		EnableTrace().
		Get(url)

	if err != nil {
		g.log.Log("GetDataHistorical - Repo - API Rest Error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return []GetDataHistoricalResponse{}, err
	}

	if resp.StatusCode() != 200 {
		g.log.Log("GetDataHistorical - Repo - Invalid reponse StatusCode", resp.StatusCode(), "response", string(resp.Body()), constants.UUID, ctx.Value(constants.UUID))
		return []GetDataHistoricalResponse{}, errors.New("Invalid ResponseCode")
	}

	if errU := json.Unmarshal(resp.Body(), &response); errU != nil {
		g.log.Log("GetDataHistorical - Repo - Unmarshal - Error", errU.Error(), constants.UUID, ctx.Value(constants.UUID))
		return []GetDataHistoricalResponse{}, errU
	}
	return response, err
}
