package telegram

import (
	"looking-for-remote-jobs/src/config"
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SendToTelegram() {
	log := zap.NewProductionEncoderConfig()
	log.EncodeLevel = zapcore.CapitalLevelEncoder
	log.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(log), os.Stdout, zap.InfoLevel))

	config.LoadEnv()
	updater, err := gotgbot.NewUpdater(logger, config.Telegram_Token)
	if err != nil {
		logger.Panic("Failed!")
		return
	}
	logger.Sugar().Info("Success!")
	updater.StartCleanPolling()
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text, echo))
	updater.Idle()
}

func echo(b ext.Bot, u *gotgbot.Update) error {
	b.SendMessage(u.EffectiveChat.Id, u.EffectiveMessage.Text)
	return nil
}
