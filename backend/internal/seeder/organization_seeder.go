package seeder

import (
	"context"
	"data-merging/internal/domain"
	"data-merging/internal/repository"
	"log"
	"os"
	"strconv"
	"time"
)

func createSeedOrganizations() []domain.Organization {
	organizationsSeedCountStr := os.Getenv("ORGANIZATIONS_SEED_COUNT")
	organizationsSeedCount, err := strconv.Atoi(organizationsSeedCountStr)
	if err != nil {
		log.Fatalf("Error converting ORGANIZATIONS_SEED_COUNT to int: %v", err)
	}

	// Create a slice to hold the organizations with the specified seed count
	organizations := make([]domain.Organization, 0, organizationsSeedCount)

	// Define some organizations
	org1 := domain.Organization{
		Base: domain.Base{
			InsertDate: time.Now(),
			UpdateDate: time.Now(),
		},
		Name: "BMW",
		Address: domain.Address{
			City:        "Novi Sad",
			Country:     "Serbia",
			AddressLine: "Rumenacki put 125",
			PostalCode:  "21000",
		},
	}

	org2 := domain.Organization{
		Base: domain.Base{
			InsertDate: time.Now(),
			UpdateDate: time.Now(),
		},
		Name: "Mercedes",
		Address: domain.Address{
			City:        "Novi Sad",
			Country:     "Serbia",
			AddressLine: "Rumenacki put 121",
			PostalCode:  "21000",
		},
	}

	org3 := domain.Organization{
		Base: domain.Base{
			InsertDate: time.Now(),
			UpdateDate: time.Now(),
		},
		Name: "Honda",
		Address: domain.Address{
			City:        "Novi Sad",
			Country:     "Serbia",
			AddressLine: "Rumenacki put 123",
			PostalCode:  "21000",
		},
	}

	// Append each organization to the slice
	organizations = append(organizations, org1, org2, org3)

	return organizations
}

// GenerateProfiles creates and saves demo profiles to MongoDB
func GenerateOrganizations(repo *repository.MongoOrganizationRepository) {
	organizations := createSeedOrganizations()

	// Use repository to interact with MongoDB
	collection := repo.Client().Database("data_merging").Collection("organizations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, organization := range organizations {
		_, err := collection.InsertOne(ctx, organization)
		if err != nil {
			log.Println("Error inserting demo organization:", err)
		}
	}
}
