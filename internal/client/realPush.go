package client

import (
	"context"
	"notification-service/internal/util"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func newRealPushClientFromConfig(credentialsFile string) *realPush {
	ctx := context.Background()

	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize Firebase app")
		return nil
	}

	c, err := app.Messaging(ctx)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("Failed to initialize Firebase Messaging app")
		return nil
	}

	return &realPush{
		client: c,
	}
}

type realPush struct {
	client *messaging.Client
}

func (rp *realPush) SendMessage(title string, body string, to string) error {
	// TODO: rp.client.SendAll()
	_, err := rp.client.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: to,
	})
	return err
}
