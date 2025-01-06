package user_management

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	firstNames = []string{"John", "Jane", "Mike", "Sarah", "David", "Emma", "James", "Emily", "Robert", "Lisa"}
	lastNames  = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}
	countries  = []string{"USA", "UK", "Canada", "Australia", "Germany", "France", "Spain", "Italy", "Japan", "Brazil"}
	currencies = []string{"USD", "GBP", "CAD", "AUD", "EUR", "JPY", "BRL"}
	teams      = []string{"Red Bull Racing", "Mercedes", "Ferrari", "McLaren", "Alpine", "AlphaTauri", "Aston Martin", "Williams", "Alfa Romeo", "Haas"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomUsers(count int) []User {
	users := make([]User, count)
	timestamp := time.Now().UnixNano()
	
	for i := 0; i < count; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		uniqueID := fmt.Sprintf("%d_%d", timestamp, i)
		username := fmt.Sprintf("%s_%s_%s", firstName, lastName, uniqueID)
		email := fmt.Sprintf("%s.%s_%s@example.com", firstName, lastName, uniqueID)
		
		phoneNumber := fmt.Sprintf("+%d%d%d", rand.Intn(99), rand.Intn(999999999), rand.Intn(999999))
		country := countries[rand.Intn(len(countries))]
		currency := currencies[rand.Intn(len(currencies))]
		team := teams[rand.Intn(len(teams))]
		profilePicURL := fmt.Sprintf("https://example.com/profiles/%s.jpg", username)

		users[i] = User{
			FullName:          fmt.Sprintf("%s %s", firstName, lastName),
			Email:             email,
			Username:          username,
			PasswordHash:      generateRandomString(32),
			DateOfBirth:       time.Now().AddDate(-rand.Intn(50)-18, -rand.Intn(12), -rand.Intn(28)),
			PhoneNumber:       &phoneNumber,
			GovernmentID:      fmt.Sprintf("ID-%s", generateRandomString(10)),
			Address:           fmt.Sprintf("%d %s St, City-%s", rand.Intn(999), generateRandomString(8), generateRandomString(6)),
			AccountStatus:     "ACTIVE",
			Role:              "USER",
			EmailVerified:     false,
			Country:           &country,
			PreferredCurrency: &currency,
			FavoriteTeam:      &team,
			ProfilePictureURL: &profilePicURL,
			RegistrationDate:  time.Now(),
		}
	}
	return users
}
