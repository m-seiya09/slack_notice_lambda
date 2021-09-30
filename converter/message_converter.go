package converter

import (
	"encoding/json"
	"fmt"
	"main/logger"
	"main/model"

	"github.com/aws/aws-lambda-go/events"
)

// slack通知する時の名前
var Username string = "ユーザー名"

// 引数で受け取ったチャンネル宛にSlackWebhookで使用するパラメータを作成する
func MakePostParameter(channel string, logEvent events.CloudwatchLogsData) (string, error) {
	slack := model.Slack{}
	slack.Channel = channel
	slack.Username = Username

	for _, logEvent := range logEvent.LogEvents {
		logger.Info(fmt.Sprintf("Log_ID = %v | Message = %v | Timestamp = %v ", logEvent.ID, logEvent.Message, logEvent.Timestamp))

		// eventの中に保存されているログ(Json)を構造体変換する
		cloudWatchMessage := model.CloudWatchMessage{}
		err := json.Unmarshal([]byte(logEvent.Message), &cloudWatchMessage)
		if err != nil {
			logger.Error(fmt.Sprintf("CloudWatchMessageの構造体変換へ失敗しました err = %v", err))

			slack.Content += fmt.Sprintf("<!here>\nログフォーマット失敗\n``` %v ```", logEvent.Message)
			json, _ := json.Marshal(slack)
			return string(json), nil
		}

		logger.Info(fmt.Sprintf("CloudWatchMessage 変換完了 model = %v \n ", cloudWatchMessage))

		// slack送信する内容をcontentに保存していく
		content := fmt.Sprintf("【発生時刻】%v\n", cloudWatchMessage.Datetime)
		content += fmt.Sprintf(
			"【環境】%v \n【ExceptionClass】\n %v \n【Message】\n %v \n【File】\n %v",
			cloudWatchMessage.Channel,
			cloudWatchMessage.Context.Exception.Class,
			cloudWatchMessage.Context.Exception.Message,
			cloudWatchMessage.Context.Exception.File,
		)

		slack.Content += fmt.Sprintf("<!here>\n```%v```", content)
	}
	json, err := json.Marshal(slack)

	return string(json), err
}
