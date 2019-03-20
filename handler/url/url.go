package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pulpfree/univsales-pdf-url/config"
	"github.com/pulpfree/univsales-pdf-url/model"
	"github.com/pulpfree/univsales-pdf-url/process"
	"github.com/pulpfree/univsales-pdf-url/validate"

	log "github.com/sirupsen/logrus"
)

var cfg *config.Config

func init() {
	cfg = &config.Config{}
	err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// Response data format
type Response struct {
	Code      int         `json:"code"`      // HTTP status code
	Data      interface{} `json:"data"`      // Data payload
	Message   string      `json:"message"`   // Error or status message
	Status    string      `json:"status"`    // Status code (error|fail|success)
	Timestamp int64       `json:"timestamp"` // Machine-readable UTC timestamp in nanoseconds since EPOCH
}

// SignedURL struct
type SignedURL struct {
	URL string `json:"url"`
}

// HandleRequest function
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var err error

	hdrs := make(map[string]string)
	hdrs["Content-Type"] = "application/json"
	t := time.Now()

	// If this is a ping test, intercept and return
	if req.HTTPMethod == "GET" {
		log.Info("Ping test in handleRequest")
		return gatewayResponse(Response{
			Code:      200,
			Data:      "pong",
			Status:    "success",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	// Check for auth header
	/* if req.Headers["Authorization"] == "" {
		return gatewayResponse(Response{
			Code:      401,
			Message:   "Missing Authorization header",
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	} */

	// Set auth config
	/* auth, err := auth.New(&auth.Config{
		ClientID:       cfg.CognitoClientID,
		PoolID:         cfg.CognitoPoolID,
		Region:         cfg.CognitoRegion,
		JwtAccessToken: req.Headers["Authorization"],
	})
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("authentication error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	} */

	// Validate JWT Token
	/* err = auth.Validate()
	if err != nil {
		return gatewayResponse(Response{
			Code:      401,
			Message:   fmt.Sprintf("token validation error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	} */

	// Set and validate request params
	var r *model.Request
	json.Unmarshal([]byte(req.Body), &r)
	err = validate.RequestInput(r)
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("request validation error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	// Process request
	process, err := process.New(r, cfg)
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("report fetch error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	url, err := process.CreateURL()
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("report create url error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}
	log.Infof("signed url created %s", url)

	return gatewayResponse(Response{
		Code:      201,
		Data:      SignedURL{URL: url},
		Status:    "success",
		Timestamp: t.Unix(),
	}, hdrs), nil
}

func main() {
	lambda.Start(HandleRequest)
}

func gatewayResponse(resp Response, hdrs map[string]string) events.APIGatewayProxyResponse {

	body, _ := json.Marshal(&resp)
	if resp.Status == "error" {
		log.Errorf("Error: status: %s, code: %d, message: %s", resp.Status, resp.Code, resp.Message)
	}

	return events.APIGatewayProxyResponse{Body: string(body), Headers: hdrs, StatusCode: resp.Code}
}
