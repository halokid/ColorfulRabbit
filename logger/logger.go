package logger

import (
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/halokid/ColorfulRabbit/file"
)

// var Logger *zap.SugaredLogger

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

func InitLogger(logLevelInit, logFolder, logFile, logExt string) {
  log.Println("-->>> Pkg logger init()")
  writeSyncer := getLogWriter(logFolder, logFile, logExt)
  encoder := getEncoder()

  logLevel := zapcore.DebugLevel
  switch logLevelInit {
  case "info":
    logLevel = zapcore.InfoLevel
  case "warn":
    logLevel = zapcore.WarnLevel
  case "error":
    logLevel = zapcore.ErrorLevel
  }

  // core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
  core := zapcore.NewCore(encoder, writeSyncer, logLevel)

  logger := zap.New(core)
  SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
  // return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
  // return zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

  config := zap.NewDevelopmentEncoderConfig()
  config.EncodeLevel = zapcore.CapitalColorLevelEncoder
  // return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
  return zapcore.NewConsoleEncoder(config)
}

func getLogWriter(logFolder, logFile, logExt string) zapcore.WriteSyncer {
  // file, _ := os.Create("./test.log")
  // file, _ := os.OpenFile("./log/runtime.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
  file, _ := os.OpenFile(logFolder + logFile + logExt, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
  ws := io.MultiWriter(file, os.Stdout) // both console and file

  go backupLog(logFolder, logFile, logExt)

  // return zapcore.AddSync(file)
  return zapcore.AddSync(ws)
}

func backupLog(logFolder, logFile, logExt string) {
  ticker := time.NewTicker( 60 * time.Second)
  for {
    select {
    case <-ticker.C:
      backupFile := time.Now().Format("2006-01-02_15-04-05")
      // copy log file
      src := logFolder + logFile + logExt
      dst := logFolder + logFile + "-" + backupFile + logExt
      file.Copy(src, dst)

      // empty log file
      err := file.Empty(src)
      if err != nil {
        log.Printf("Logger backupLog err -->>> %+v", err)
      }
    }
  }
}
