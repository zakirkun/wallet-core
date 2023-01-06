package migration

import (
	"github.com/walletkita/wallet-core/app/domain/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.Migrator().CreateTable(&models.User{}, &models.AccountInformation{}, &models.WalletCard{}, &models.WalletTag{}, &models.Security{}, &models.Notification{}, &models.TransactionHistory{})

	// Fk do Details
	db.Migrator().CreateConstraint(&models.User{}, "AccountInformation")
	db.Migrator().CreateConstraint(&models.User{}, "WalletCard")
	db.Migrator().CreateConstraint(&models.User{}, "Security")
	db.Migrator().CreateConstraint(&models.User{}, "Notification")

	db.Migrator().CreateConstraint(&models.WalletTag{}, "WalletCard")
	db.Migrator().CreateConstraint(&models.TransactionHistory{}, "User")

	return err
}
