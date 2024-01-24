package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/modules/mhsbiodata/entity"
	"tracer-study-grpc/pb"
)

type MhsBiodataHandler struct {
	pb.UnimplementedMhsBiodataServiceServer
	config config.Config
}

func NewMhsBiodataHandler(config config.Config) *MhsBiodataHandler {
	return &MhsBiodataHandler{
		config: config,
	}
}

func (mbh *MhsBiodataHandler) FetchMhsBiodataByNim(ctx context.Context, req *pb.MhsBiodataRequest) (*pb.MhsBiodataResponse, error) {
	nim := req.GetNim()

	payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	apiUrl := mbh.config.SIAK_API.URL
	apiKey := mbh.config.SIAK_API.KEY
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

	var mhsBiodata *pb.MhsBiodata
	mhsBiodata = entity.ConvertEntityToProto(&apiResponse[0])

	code := uint32(http.StatusOK)
	message := "get mhs biodata success"

	return &pb.MhsBiodataResponse{
		Code:    code,
		Message: message,
		Data:    mhsBiodata,
	}, nil
}
