package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type FCMRepository struct {
	once   sync.Once
	client *messaging.Client
	err    error
}

func NewFCMRepository() *FCMRepository {
	return &FCMRepository{}
}

func (f *FCMRepository) initClient() {
	f.once.Do(func() {
		credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_FILE")
		if credentialsPath == "" {
			f.err = errors.New("FIREBASE_CREDENTIALS_FILE no está configurado")
			return
		}

		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialsPath))
		if err != nil {
			f.err = fmt.Errorf("error inicializando Firebase app: %w", err)
			return
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			f.err = fmt.Errorf("error inicializando cliente de Messaging: %w", err)
			return
		}
		f.client = client
	})
}

func (f *FCMRepository) SendToToken(token, title, body string, data map[string]string) error {
	f.initClient()
	if f.err != nil {
		return f.err
	}
	if f.client == nil {
		return errors.New("cliente FCM no disponible")
	}

	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}

	if channelID, ok := data["android_channel_id"]; ok && channelID != "" {
		msg.Android.Notification = &messaging.AndroidNotification{
			ChannelID: channelID,
		}
	}

	_, err := f.client.Send(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("error enviando push FCM: %w", err)
	}
	return nil
}

func (f *FCMRepository) SendToTokens(tokens []string, title, body string, data map[string]string) error {
	f.initClient()
	if f.err != nil {
		return f.err
	}
	if f.client == nil {
		return errors.New("cliente FCM no disponible")
	}
	if len(tokens) == 0 {
		return errors.New("no hay tokens para enviar")
	}

	msg := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}
	if channelID, ok := data["android_channel_id"]; ok && channelID != "" {
		msg.Android.Notification = &messaging.AndroidNotification{
			ChannelID: channelID,
		}
	}

	res, err := f.client.SendEachForMulticast(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("error enviando push multicast FCM: %w", err)
	}
	if res.FailureCount == len(tokens) {
		return fmt.Errorf("falló el envío a todos los dispositivos: %d errores", res.FailureCount)
	}
	return nil
}
