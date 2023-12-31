package utils

import (
	"fmt"
	"math/rand"
	"rapigo/internal/service"
	"time"
)

func GenerateUniqId(firstName string, lastName string) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(9000) + 1000

	firstLetter := string(firstName[0])
	lastLetter := string(lastName[0])

	adminId := fmt.Sprintf("%s%s%d", firstLetter, lastLetter, randomNumber)

	response, _ := service.GetFindMyField("adminId", adminId)

	// Check if the generated adminId already exists
	if response.AdminId == adminId {
		// If it exists, generate a new one (recursive call)
		return GenerateUniqId(firstName, lastName)
	}

	return adminId
}
