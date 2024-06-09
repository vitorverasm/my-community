package supabase

import (
	"log"

	"github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/config"
)

func InitializeClient() *supabase.Client {
	env := config.LoadEnvVariables()
	client, err := supabase.NewClient(env.SupabaseUrl, env.SupabaseApiKey, nil)
	if err != nil {
		log.Println("cannot initialize client", err)
	}

	return client
}
