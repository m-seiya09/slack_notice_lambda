package logger

import "log"

var prefix string = "[Lambda Log Watch Function] "

// 初期化
func init() {
	log.SetPrefix(prefix)
}

// infoログを出力する
func Info(message string) {
	log.Println(log.Prefix() + "(INFO) " + message)
}

// errorログを出力する
func Error(message string) {
	log.Println(log.Prefix() + "(ERROR) " + message)
}
