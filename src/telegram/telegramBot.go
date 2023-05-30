package telegram

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SendToTelegram() {
	log := zap.NewProductionEncoderConfig()
	log.EncodeLevel = zapcore.CapitalLevelEncoder
	log.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(log), os.Stdout, zap.InfoLevel))

	updater, err := gotgbot.NewUpdater(logger, "Token")
	if err != nil {
		logger.Panic("Failed!")
		return
	}
	logger.Sugar().Info("Success!")
	updater.StartCleanPolling()
	updater.Idle()
}
