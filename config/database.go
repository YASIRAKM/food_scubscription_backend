package config

import (
	"log"
	"myapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=9qasp5v56q8ckkf5dc.leapcellpool.com port=6438 user=woqxphhbfkgatubbzybg password=qwcklcsidntqmlqiudevbdnatbblsf dbname=vgnnxckjzkrdzssaldva sslmode=require"
		var err error
		DB,err  = gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database")
		}
		DB.AutoMigrate(&models.User{}, &models.Subscription{}, &models.Food{})
		log.Println("Database connected and migrated successfully.")
}