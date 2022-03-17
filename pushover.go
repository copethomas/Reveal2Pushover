package main

import (
	"github.com/gregdel/pushover"
)

func SendPushNotification(msg string) error {
	app := pushover.New(conf.PushOverAppKey)
	recipient := pushover.NewRecipient(conf.PushOverUserKey)
	message := pushover.NewMessage(msg)
	_, err := app.SendMessage(message, recipient)
	return err
}
