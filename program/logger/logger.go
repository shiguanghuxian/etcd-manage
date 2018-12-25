package logger

import (
	"os"
	"strings"
	"time"

	"github.com/shiguanghuxian/etcd-manage/program/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志对象
var (
	Log *zap.SugaredLogger
)

// InitLogger 日志初始化，用于记录操作日志
func InitLogger(logPath string, isDebug bool) (*zap.SugaredLogger, error) {
	infoLogPath := ""
	errorLogPath := ""
	if logPath == "" {
		logRoot := common.GetRootDir() + "logs/"
		if isExt, _ := common.PathExists(logRoot); isExt == false {
			os.MkdirAll(logRoot, os.ModePerm)
		}
		infoLogPath = logRoot + time.Now().Format("20060102_info") + ".log"
		errorLogPath = logRoot + time.Now().Format("20060102_error") + ".log"
	} else {
		logPath = strings.TrimRight(logPath, string(os.PathSeparator))
		infoLogPath = logPath + string(os.PathSeparator) + time.Now().Format("20060102_info") + ".log"
		errorLogPath = logPath + string(os.PathSeparator) + time.Now().Format("20060102_error") + ".log"
	}

	cfg := &zap.Config{
		Encoding: "json",
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	atom := zap.NewAtomicLevel()
	if isDebug == true {
		atom.SetLevel(zapcore.DebugLevel)
		cfg.OutputPaths = []string{"stdout"}
		cfg.ErrorOutputPaths = []string{"stdout"}
	} else {
		atom.SetLevel(zapcore.InfoLevel)
		cfg.OutputPaths = []string{infoLogPath}
		cfg.ErrorOutputPaths = []string{errorLogPath}
	}
	cfg.Level = atom
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	Log = logger.Sugar()
	return Log, nil
}
