package logger

import (
	"io"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

/*
func init() {
  log.Printf("-->>> LoggerInit")
  config := zap.NewDevelopmentConfig()
  // config := zap.NewProductionConfig()
  config.EncoderConfig.TimeKey = "timestamp"
  config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info
  level := zapcore.DebugLevel
  // level := zapcore.InfoLevel
  // level := zapcore.ErrorLevel
  logLevel := zap.NewAtomicLevelAt(level) // TODO: set the log level to `warn`
  config.Level = logLevel
  config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
  // config.OutputPaths = []string{"stdout", "runtime.log"}

  logger, err := config.Build()
  if err != nil {
    log.Fatal(err)
  }

  Logger = logger.Sugar()
}
*/

var SugarLogger *zap.SugaredLogger

func InitLogger() {
    log.Println("-->>> Pkg logger init()")
    writeSyncer := getLogWriter()
    encoder := getEncoder()
    core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

    logger := zap.New(core)
    SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
    // return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    // return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

    config :=  zap.NewDevelopmentEncoderConfig()
    config.EncodeLevel = zapcore.CapitalColorLevelEncoder
    // return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
    return zapcore.NewConsoleEncoder(config)
}

func getLogWriter() zapcore.WriteSyncer {
    // file, _ := os.Create("./test.log")
    file, _ := os.OpenFile("./runtime.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
    ws := io.MultiWriter(file, os.Stdout)   // both console and file
    // return zapcore.AddSync(file)
    return zapcore.AddSync(ws)
}

func simpleHttpGet(url string) {
    SugarLogger.Debugf("Trying to hit GET request for %s", url)
    resp, err := http.Get(url)
    if err != nil {
        SugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
    } else {
        SugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
        resp.Body.Close()
    }
}




