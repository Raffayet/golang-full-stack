package main

import (
	"context"
	"data-merging/internal/api"
	"data-merging/internal/config"
	"data-merging/internal/repository"
	"data-merging/internal/seeder"
	"data-merging/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	client := config.SetupMongoDb()

	// Defer disconnecting the client until the application is shutting down
	defer client.Disconnect(context.Background())

	// Create repository, service, and handler
	datasetRepo := repository.NewDatasetRepository(client)
	datasetService := service.NewDatasetService(datasetRepo)
	datasetHandler := api.NewDatasetHandler(datasetService)

	organizationRepo := repository.NewOrganizationRepository(client)

	router := config.SetupRouter(datasetHandler)
	seeder.GenerateDemoData(datasetRepo, organizationRepo)

	startServer(router)
}

func startServer(router *mux.Router) {
	cors := config.SetupCorsPolicy()

	// Get API URI from environment variable
	apiAddress := os.Getenv("API_URI_SHORTENED")
	if apiAddress == "" {
		log.Println("API_URI not set, using default :8000")
		apiAddress = ":8000"
	}

	// Start the HTTP server with CORS middleware
	log.Printf("Server is running on %s", apiAddress)
	err := http.ListenAndServe(apiAddress, cors(router))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
