package seeder

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/domain"
	"github.com/Raffayet/data-merging/backend/internal/repository"
	"golang.org/x/exp/rand"
)

var firstNamesPool = []string{
	"John", "Jane", "Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank",
	"Ivy", "Jack", "Kathy", "Liam", "Mia", "Nina", "Oscar", "Paul", "Quinn", "Rose",
	"Sam", "Tina", "Uma", "Victor", "Wendy", "Xander", "Yara", "Zack", "Abby", "Ben",
}

var lastNamesPool = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
	"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
}

var addressPool = []string{
	"123 Main St, Springfield",
	"456 Oak Dr, Metropolis",
	"789 Elm St, Gotham City",
	"101 Maple Ave, Star City",
	"202 Pine St, Central City",
	"303 Cedar Blvd, Coast City",
	"404 Birch Ln, Smallville",
	"505 Willow Rd, Midway City",
}

func createSeedProfiles() []domain.Profile {
	profilesSeedCountStr := os.Getenv("PROFILES_SEED_COUNT")
	profilesSeedCount, err := strconv.Atoi(profilesSeedCountStr)
	if err != nil {
		log.Fatalf("Error converting PROFILES_SEED_COUNT to int: %v", err)
	}

	profiles := make([]domain.Profile, profilesSeedCount)

	for i := 0; i < profilesSeedCount; i++ {
		firstName := firstNamesPool[rand.Intn(len(firstNamesPool))]
		lastName := lastNamesPool[rand.Intn(len(lastNamesPool))]
		phoneNumber := "061220113" + strconv.Itoa(i)

		p := domain.Profile{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       strings.ToLower(firstName) + "." + strings.ToLower(lastName) + "@raffayet.com",
			PhoneNumber: phoneNumber,
			Address:     addressPool[rand.Intn(len(addressPool))],
		}
		profiles[i] = p
	}

	return profiles
}

// GenerateProfiles creates and saves demo profiles to MongoDB
func GenerateProfiles(repo *repository.MongoProfileRepository) {
	profiles := createSeedProfiles()

	// Use repository to interact with MongoDB
	collection := repo.Client().Database("data_merging").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, profile := range profiles {
		_, err := collection.InsertOne(ctx, profile)
		if err != nil {
			log.Println("Error inserting demo profile:", err)
		}
	}
}
