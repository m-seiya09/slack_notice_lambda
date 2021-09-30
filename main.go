package main

import (
	"bytes"
	"context"
	"fmt"
	"main/converter"
	"main/logger"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	webhookUrl   string = "***************"
	slackChannel        = "********"
)

func main() {
	logger.Info("lambda関数実行")

	lambda.Start(HandleRequest)

	logger.Info("lambda関数終了")
}

// 実行関数
func HandleRequest(context context.Context, LogEvent events.CloudwatchLogsEvent) {
	data, _ := LogEvent.AWSLogs.Parse()

	logger.Info("MakePostParameter 実行")
	param, err := converter.MakePostParameter(slackChannel, data)
	logger.Info("MakePostParameter 終了")

	if err != nil {
		logger.Error(fmt.Sprintf("MakePostParameter error = %v \n", err))
		panic(err)
	}

	sendSlack(param)
}

// webhookを使用してslackに送信する
func sendSlack(json string) {
	logger.Info("sendSlack関数実行")
	request, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer([]byte(json)))

	if err != nil {
		logger.Error(fmt.Sprintf("new request error = %v \n", err))
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	responce, err := client.Do(request)

	if err != nil {
		logger.Error(fmt.Sprintf("slack post error %v \n", err))
		panic(err)
	}
	defer responce.Body.Close()

	logger.Info("sendSlack関数終了")
}
