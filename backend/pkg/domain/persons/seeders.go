package persons

import (
	"time"

	"gorm.io/gorm"
)

func SeedPerson(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {

		person := Person{
			FirstName: "Oğuz",
			LastName:  "Sarıçam",
			Gsm:       "5372112339",
			Email:     "",
			BirthDate: time.Now(),
		}
		person.SetDefaultsviaCreation()
		if err := tx.Create(&person).Error; err != nil {
			return err
		}

		return nil
	})
}
