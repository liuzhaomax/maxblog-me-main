package main

import (
	"context"
	"maxblog-me-main/internal/app"
	"maxblog-me-main/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
