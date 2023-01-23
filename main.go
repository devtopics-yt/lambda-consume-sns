package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger
var lambdaName string

func init() {
	l, _ := zap.NewProduction()
	logger = l

	lambdaName = os.Getenv("LAMBDA_NAME")
}

type Event struct {
	OrderID string `json:"orderId"`
	TS      int    `json:"ts"`
}

func MyHandler(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		logger.Info("received sns event", zap.Any("record", record), zap.Any("lambdaName", lambdaName))

		// Decode JSON
		event := &Event{}
		err := json.Unmarshal([]byte(record.SNS.Message), event)
		if err != nil {
			return err
		}

		logger.Info("decoded event", zap.Any("event", event), zap.Any("lambdaName", lambdaName))
	}

	return nil
}

func main() {
	lambda.Start(MyHandler)
}
