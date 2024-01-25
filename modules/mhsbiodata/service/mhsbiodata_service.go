package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/entity"
)

type MhsBiodataService struct {
	cfg config.Config
}

type MhsBiodataServiceUseCase interface {
	FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodata, error)
}

func NewMhsBiodataService(cfg config.Config) *MhsBiodataService {
	return &MhsBiodataService{
		cfg: cfg,
	}
}

func (svc *MhsBiodataService) FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodata, error) {
	payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	apiUrl := svc.cfg.SIAK_API.URL
	apiKey := svc.cfg.SIAK_API.KEY
	reqHttp, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	reqHttp.Header.Set("Api-Key", apiKey)
	reqHttp.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqHttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse []entity.MhsBiodata
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse[0], nil
}