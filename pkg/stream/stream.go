package stream

import (
	"log"
	"time"

	streamSDK "github.com/GetStream/stream-chat-go/v5"
)

func InitializeClient(apiKey string, apiSecret string) *streamSDK.Client {
	client, err := streamSDK.NewClient(apiKey, apiSecret)
	if err != nil {
		log.Println("cannot initialize stream client", err)
	}

	return client
}

type StreamCommunicationProvider struct {
	Client *streamSDK.Client
}

func (stream *StreamCommunicationProvider) GetUserCommunicationToken(userEmail string) (userCommunicationToken string, err error) {
	token, error := stream.Client.CreateToken(userEmail, time.Time{})

	if error != nil {
		return "", error
	}

	return token, nil
}
