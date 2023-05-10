package database

import (
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) {
	var errs []error
	err := runCategoriesSeeder(db)
	errs = append(errs, err)
	for _, e := range errs {
		if e != nil {
			log.Errorf("seeder error-> %v", e)
		}
	}

}

func runCategoriesSeeder(db *gorm.DB) error {
	categories := []models.Category{
		{Name: "Shopping", Custom: false, Icon: "🛍️", Unicode: "U+1F6CD"},
		{Name: "Utilities and Bills", Custom: false, Icon: "", Unicode: ""},
		{Name: "Housing", Custom: false, Icon: "", Unicode: ""},
		{Name: "Entertainment", Custom: false, Icon: "", Unicode: ""},
		{Name: "Travel", Custom: false, Icon: "", Unicode: ""},
		{Name: "Transportation", Custom: false, Icon: "", Unicode: ""},
		{Name: "Income", Custom: false, Icon: "", Unicode: ""},
		{Name: "Investment", Custom: false, Icon: "", Unicode: ""},
		{Name: "Phone & Internet", Custom: false, Icon: "", Unicode: ""},
		{Name: "Food", Custom: false, Icon: "", Unicode: ""},
		{Name: "Healthcare", Custom: false, Icon: "", Unicode: ""},
		{Name: "Loan Repayment", Custom: false, Icon: "", Unicode: ""},
		{Name: "Loan Out", Custom: false, Icon: "", Unicode: ""},
		{Name: "Transfer", Custom: false, Icon: "", Unicode: ""},
		{Name: "Online Transaction", Custom: false, Icon: "", Unicode: ""},
		{Name: "Offline Transaction", Custom: false, Icon: "", Unicode: ""},
		{Name: "Bank Charges", Custom: false, Icon: "", Unicode: ""},
		{Name: "ATM Withdrawal", Custom: false, Icon: "", Unicode: ""},
		{Name: "Miscellaneous", Custom: false, Icon: "", Unicode: ""},
		{Name: "Gifts & Donations", Custom: false, Icon: "", Unicode: ""},
		{Name: "Education", Custom: false, Icon: "", Unicode: ""},
	}
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories).Error; err != nil {
		return err
	}

	return nil
}
