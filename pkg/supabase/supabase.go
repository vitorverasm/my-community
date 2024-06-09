package supabase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
	"github.com/vitorverasm/my-community/config"
)

func InitializeClient() *supabase.Client {
	env := config.LoadEnvVariables()
	client, err := supabase.NewClient(env.SupabaseUrl, env.SupabaseApiKey, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	return client
}
