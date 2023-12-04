package view

import (
	"fmt"

	"github.com/Devil666face/fiber/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Map fiber.Map

func (m Map) get(key string) any {
	if val, ok := m[key]; ok {
		return val
	}
	return fmt.Errorf("not found value for key: %s in view map", key)
}

func (m Map) getUser() models.User {
	if user, ok := m.get(UserKey).(models.User); ok {
		return user
	}
	return models.User{}
}

func (m Map) notUser() bool {
	return (m.getUser() == models.User{})
}
