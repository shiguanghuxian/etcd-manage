package logger

import (
	"net/url"
	"os"
	"runtime"
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
	// errorLogPath := ""
	if logPath == "" {
		logRoot := common.GetRootDir() + "logs" + string(os.PathSeparator)
		if isExt, _ := common.PathExists(logRoot); isExt == false {
			os.MkdirAll(logRoot, os.ModePerm)
		}
		infoLogPath = logRoot + time.Now().Format("20060102") + ".log"
		// errorLogPath = logRoot + time.Now().Format("20060102_error") + ".log"
	} else {
		logPath = strings.TrimRight(logPath, string(os.PathSeparator))
		infoLogPath = logPath + string(os.PathSeparator) + time.Now().Format("20060102") + ".log"
		// errorLogPath = logPath + string(os.PathSeparator) + time.Now().Format("20060102_error") + ".log"
	}

	// 兼容win根完整路径问题
	zap.RegisterSink("winfile", func(u *url.URL) (zap.Sink, error) {
		return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	})

	cfg := &zap.Config{
		Encoding: "json",
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	atom := zap.NewAtomicLevel()
	if isDebug == true {
		atom.SetLevel(zapcore.DebugLevel)
		cfg.OutputPaths = []string{"stdout"}
		// cfg.ErrorOutputPaths = []string{"stdout"}
	} else {
		atom.SetLevel(zapcore.InfoLevel)
		if runtime.GOOS == "windows" {
			cfg.OutputPaths = []string{"winfile:///" + infoLogPath}
		} else {
			cfg.OutputPaths = []string{infoLogPath}
		}
		// cfg.ErrorOutputPaths = []string{errorLogPath}
	}
	cfg.Level = atom
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	Log = logger.Sugar()
	return Log, nil
}
