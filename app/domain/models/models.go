package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserID        int                `db:"user_id" gorm:"primaryKey"`
	FullName      string             `db:"full_name"`
	Email         string             `db:"email"`
	PhoneNumber   string             `db:"phone"`
	Password      string             `db:"password"`
	LastIPLogin   string             `db:"last_ip_login"`
	LastLoginTime time.Time          `db:"last_login_time"`
	IsLocked      bool               `db:"is_locked"`
	Pin           string             `db:"pin"`
	Details       AccountInformation `gorm:"references:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Wallets       []WalletCard       `gorm:"references:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Security      []Security         `gorm:"references:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Notification  []Notification     `gorm:"references:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AccountInformation struct {
	gorm.Model

	IdDetail  int    `db:"id_detail" gorm:"primaryKey"`
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Address   string `db:"address"`
	ZipCode   string `db:"zip_code"`
}

type WalletCard struct {
	gorm.Model

	WalletID     int    `db:"wallet_id" gorm:"primaryKey"`
	UsersId      int    `db:"user_id"`
	WalletNumber string `db:"wallet_number"`
	WalletTag    int    `db:"wallet_tag"`
	Balance      int64  `db:"balance"`
	IsLock       bool   `db:"is_lock"`
}

type WalletTag struct {
	gorm.Model

	TagId   int          `db:"tag_id" gorm:"primaryKey"`
	UserId  int          `db:"user_id"`
	TagName string       `db:"tag_name"`
	Wallet  []WalletCard `gorm:"references:WalletTag;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Security struct {
	gorm.Model

	Id         int    `db:"id" gorm:"primaryKey"`
	UsersID    int    `db:"user_id"`
	Type       string `db:"type"`
	Content    string `db:"content"`
	ActionNeed string `db:"action_need"`
	Ip         string `db:"ip"`
}

type Notification struct {
	gorm.Model

	TagId   int    `db:"tag_id" gorm:"primaryKey"`
	UserId  int    `db:"user_id"`
	Title   string `db:"title"`
	Subject string `db:"subject"`
	Message string `db:"message"`
	IsOpen  bool   `db:"isOpen"`
	Sender  string `db:"sender"`
}

type TransactionHistory struct {
	gorm.Model

	LogID   int    `db:"log_id" gorm:"primaryKey"`
	UserId  int    `db:"user_id"`
	LogType string `db:"log_type"`
	Message string `db:"message"`
	Amount  string `db:"amount"`
	Status  string `db:"status"`
	User    User   `gorm:"references:UserId"`
}
