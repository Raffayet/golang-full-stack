package seeder

import (
	"context"
	"data-merging/internal/repository"
	"log"
	"time"
)

func GenerateDemoData(datasetRepo *repository.MongoDatasetRepository, organizationRepo *repository.MongoOrganizationRepository) {
	// Create demo data
	CleanDemoData(datasetRepo)
	GenerateOrganizations(organizationRepo)
	GenerateSeedDatasets(datasetRepo)
}

// CleanDemoData removes the demo profiles from MongoDB
func CleanDemoData(repo *repository.MongoDatasetRepository) {
	// Use the MongoDB client to drop the entire database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.Client().Database("data_merging").Drop(ctx)
	if err != nil {
		log.Println("Error cleaning the database:", err)
	} else {
		log.Println("Database cleaned successfully, all collections dropped.")
	}
}
