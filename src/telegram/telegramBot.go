package telegram

import (
	"fmt"
	"looking-for-remote-jobs/src/config"
	"looking-for-remote-jobs/src/service"
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SendMessageToTelegram() {
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
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text, sendToTelegram))
	updater.Idle()
}

func receiveMessage(message string) string {
	text := message
	fmt.Println("MESSAGE", text)
	return service.GetAllOportunities(text)
}

func sendToTelegram(b ext.Bot, u *gotgbot.Update) error {
	text := u.EffectiveMessage.Text
	if text != "" {
		b.SendMessage(u.EffectiveChat.Id, receiveMessage(text))
	} else {
		b.SendMessage(u.EffectiveChat.Id, "Enter a job")
	}
	return nil
}
