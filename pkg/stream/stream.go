package stream

import (
	"log"
	"time"

	streamSDK "github.com/GetStream/stream-chat-go/v5"
	"github.com/vitorverasm/my-community/config"
)

func InitializeClient() *streamSDK.Client {
	env := config.LoadEnvVariables()
	client, err := streamSDK.NewClient(env.StreamApiKey, env.StreamApiSecret)
	if err != nil {
		log.Println("cannot initialize stream client", err)
	}

	return client
}

type StreamCommunicationProvider struct {
	Client *streamSDK.Client
}

func (stream *StreamCommunicationProvider) GetUserInteractionToken(userEmail string) (userInteractionToken string, err error) {
	token, error := stream.Client.CreateToken(userEmail, time.Time{})

	if error != nil {
		return "", error
	}

	return token, nil
}
