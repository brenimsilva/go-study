package initializers

import "brenimsilva/auth/models"

func Migration() {
	DB.AutoMigrate(&models.User{})
}
