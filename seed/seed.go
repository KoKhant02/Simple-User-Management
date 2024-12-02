package seed

import (
	"log"
	"user-crud/models"
	"user-crud/utils"

	"gorm.io/gorm"
)

func SeedMasterAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Where("email = ?", "admin@mail.com").Count(&count)
	if count == 0 {
		hashedPassword, _ := utils.HashPassword("admin123")
		admin := models.User{
			Name:     "Master Admin",
			Email:    "admin@mail.com",
			Password: hashedPassword,
			Role:     "admin",
		}
		if err := db.Create(&admin).Error; err != nil {
			log.Println("Failed to seed master admin:", err)
		} else {
			log.Println("Master admin seeded successfully")
		}
	}
}
