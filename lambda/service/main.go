package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pennsieve/direct-download-metrics/service/handler"
)

func main() {
	lambda.Start(handler.DirectDownloadMetricsHandler)
}
