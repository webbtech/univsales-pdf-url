package main

import (
	"encoding/json"
	"time"

	pres "github.com/pulpfree/lambda-go-proxy-response"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/epsagon/epsagon-go/epsagon"
	"github.com/pulpfree/univsales-pdf-url/config"
	"github.com/pulpfree/univsales-pdf-url/model"
	"github.com/pulpfree/univsales-pdf-url/process"
	"github.com/pulpfree/univsales-pdf-url/validate"
)

// SignedURL struct
type SignedURL struct {
	URL string `json:"url"`
}

const (
	epsagonAppName = "univsales"
	epsagonToken   = "73993039-d583-43ad-84eb-1a443e257274"
)

var cfg *config.Config

func init() {
	cfg = &config.Config{}
	err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// HandleRequest function
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var err error

	hdrs := make(map[string]string)
	hdrs["Content-Type"] = "application/json"
	hdrs["Access-Control-Allow-Origin"] = "*"
	hdrs["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,PUT"
	hdrs["Access-Control-Allow-Headers"] = "Authorization,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"

	if req.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{Body: string("null"), Headers: hdrs, StatusCode: 200}, nil
	}

	t := time.Now()

	// If this is a ping test, intercept and return
	if req.HTTPMethod == "GET" {
		log.Info("Ping test in handleRequest")
		return pres.ProxyRes(pres.Response{
			Code:      200,
			Data:      "pong",
			Status:    "success",
			Timestamp: t.Unix(),
		}, hdrs, nil), nil
	}

	// Set and validate request params
	var r *model.Request
	json.Unmarshal([]byte(req.Body), &r)

	// validate input
	err = validate.RequestInput(r)
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	// Process request
	process, err := process.New(r, cfg)
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	url, err := process.CreateURL()
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	urlStr := url[0:100]
	log.Infof("signed url created: %s", urlStr)

	return pres.ProxyRes(pres.Response{
		Code:      201,
		Data:      SignedURL{URL: url},
		Status:    "success",
		Timestamp: t.Unix(),
	}, hdrs, nil), nil
}

func main() {
	log.Println("enter main")
	lambda.Start(epsagon.WrapLambdaHandler(
		epsagon.NewTracerConfig(epsagonAppName, epsagonToken),
		HandleRequest))
}
