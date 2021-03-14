package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logfile, err := os.Create("./learn_log/logrus_related/log.txt")
	if err != nil {
		logrus.Errorf("file open error: %s", err)
	}
	defer logfile.Close()
	// 设置日志的输出文件
	logrus.SetOutput(logfile)
	// 在日志中携带方法名，会有一定的性能开销
	logrus.SetReportCaller(true)
	// 以 JSON 格式输出日志
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Hello world")
	
	// logrus 推荐进行结构化记日志，通用字段的记录：
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common filed",
	})
	contextLogger.Info("this is logger for common info")
	// 类似 Failed to send event %s to topic %s with key %d" 这样的信息应该改为：
	// 	logrus.WithFields(logrus.Fields{
	//  "event": event,
	//  "topic": topic,
	//  "key": key,
	// }).Fatal("Failed to send event")
}

