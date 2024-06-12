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

	environment := os.Getenv("APP_ENV")
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseApiKey := os.Getenv("SUPABASE_API_KEY")
	streamApiKey := os.Getenv("STREAM_API_KEY")
	streamApiSecret := os.Getenv("STREAM_API_SECRET")

	if supabaseUrl == "" {
		log.Fatal("SUPABASE_URL not found in .env file")
	}

	if supabaseApiKey == "" {
		log.Fatal("SUPABASE_API_KEY not found in .env file")
	}

	if streamApiKey == "" {
		log.Fatal("STREAM_API_KEY not found in .env file")
	}

	if streamApiSecret == "" {
		log.Fatal("STREAM_API_SECRET not found in .env file")
	}

	if environment == "" {
		log.Fatal("APP_ENV not found in .env file")
	}

	env := types.ApplicationEnv{
		Environment:     environment,
		SupabaseUrl:     supabaseUrl,
		SupabaseApiKey:  supabaseApiKey,
		StreamApiKey:    streamApiKey,
		StreamApiSecret: streamApiSecret,
	}

	return env
}
