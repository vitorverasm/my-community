package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vitorverasm/my-community/config"
	"github.com/vitorverasm/my-community/handlers"
	"github.com/vitorverasm/my-community/pkg/stream"
	"github.com/vitorverasm/my-community/pkg/supabase"
)

func main() {
	env := config.LoadEnvVariables()

	if env.Environment == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	supabaseAuthProvider := &supabase.SupabaseAuthProvider{
		Client: supabase.InitializeClient(env.SupabaseUrl, env.StreamApiKey),
	}

	streamCommunicationProvider := &stream.StreamCommunicationProvider{
		Client: stream.InitializeClient(env.StreamApiKey, env.StreamApiSecret),
	}

	r.POST("/login", func(c *gin.Context) {
		handlers.NewLoginHandler(supabaseAuthProvider)(c)
	})

	r.POST("/register", func(c *gin.Context) {
		handlers.NewSignUpHandler(supabaseAuthProvider, streamCommunicationProvider)(c)
	})

	r.Run(":3000")
	log.Println("Running server on port 3000...")
}
