package main

import (
	"looking-for-remote-jobs/src/service"
	//"looking-for-remote-jobs/src/telegram"
)

func main() {
	//telegram.SendMessageToTelegram()

	teste := ".net developer"
	service.GetAllOportunities(teste)
}
