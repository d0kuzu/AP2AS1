package main

import (
	"assign1/internal/api"
	appModule "assign1/internal/app"
	"assign1/internal/cleanup"
	"assign1/internal/config"
	"assign1/internal/database"
	"assign1/internal/database/repositories"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.InitPool(ctx, settings)
	if err != nil {
		log.Fatal("Failed to initialize database pool:", err)
	}

	cm := &cleanup.CleanupManager{}
	cm.Add(pool.Close)
	go cm.Wait()

	repos := repositories.NewRepositories(pool)

	app := appModule.NewApp(repos, settings)

	api.RouterStart(app)
}
