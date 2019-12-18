package log

import (
	"go.uber.org/zap"
	"sync"
)

var (
	once sync.Once
	ZapLog *zap.SugaredLogger
)

func InitLog() {
	once.Do(func() {
		//config := zap.Config{
		//	Encoding:			"json",
		//	OutputPaths:       []string{"./runtime/log/error.log", "stdout"},
		//	ErrorOutputPaths:  []string{"./runtime/log/zaperror.log", "stderr"},
		//}
		//logger, err := config.Build()
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		ZapLog = logger.Sugar()
	})
}