package seed

import (
	"log"
	"time"

	"github.com/aniruddha2000/admybrand-assignment/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Username: "test1",
		Email:    "test@test1.com",
		Password: "1234",
		DOB:      time.Date(2000, time.Month(11), 9, 0, 0, 0, 0, time.UTC),
		Address: "294, SV Road",
		Description: "Hey It's the description",
	},
	{
		Username: "test2",
		Email:    "test@test2.com",
		Password: "1234",
		DOB:      time.Date(2000, time.Month(10), 14, 0, 0, 0, 0, time.UTC),
		Address: "39 SG Road",
		Description: "Hey It's the description",
	},
}

// Load the above struct the for the first time in the database
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}
	}
}
