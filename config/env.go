package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vitorverasm/my-community/types"
)

func LoadEnvVariables() types.ApplicationEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseApiKey := os.Getenv("SUPABASE_API_KEY")

	if supabaseUrl == "" || supabaseApiKey == "" {
		log.Fatal("SUPABASE_URL or SUPABASE_API_KEY not found in .env file")
	}

	env := types.ApplicationEnv{
		SupabaseUrl:    supabaseUrl,
		SupabaseApiKey: supabaseApiKey,
	}

	return env
}
