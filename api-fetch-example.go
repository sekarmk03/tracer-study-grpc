package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"your_module/your_protobuf_package"
)

type server struct{}

func (s *server) FetchData(ctx context.Context, req *your_protobuf_package.Request) (*your_protobuf_package.Response, error) {
	// Extract the value of "nim" from the request
	nim := req.GetNim()

	// Create a JSON payload
	payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Make an HTTP request to the external API
	apiURL := "https://siak.upi.edu/api/siak/mhs_biodata"
	reqHTTP, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	// Set the API-Key header
	reqHTTP.Header.Set("Api-Key", "your_api_key")

	// Set Content-Type header
	reqHTTP.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(reqHTTP)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a gRPC response
	response := &your_protobuf_package.Response{
		StatusCode:  int32(resp.StatusCode),
		ApiResponse: string(body),
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	your_protobuf_package.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
