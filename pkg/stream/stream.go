package stream

import (
	"log"
	"time"

	stream "github.com/GetStream/stream-chat-go/v5"
	"github.com/vitorverasm/my-community/config"
)

func GetToken(userId string) (string, error) {
	env := config.LoadEnvVariables()
	client, err := stream.NewClient(env.StreamApiKey, env.StreamApiSecret)

	if err != nil {
		log.Println("Failed to initialize stream client - Error: " + err.Error())
	}

	return client.CreateToken(userId, time.Time{})
}
