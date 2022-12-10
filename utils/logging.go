package utils

import (
	"io"
	"log"
	"os"
)

/*
このファイルを作ることで、
1. 読み書き
2. ファイルの作成
3. ファイルの追記
4. パーミッションを0666に設定
のログを出力することができる
*/
func LoggingSetting(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
