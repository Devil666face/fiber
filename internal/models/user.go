package models

import (
	"errors"

	"github.com/Devil666face/fiber/internal/web/validators"
	"github.com/Devil666face/fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	ErrPasswordEncrypt = fiber.ErrInternalServerError
	ErrUserNotUniq     = fiber.NewError(fiber.StatusBadRequest, "User already create")
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null" form:"email" validate:"required,email"`
	Password        string `gorm:"not null" form:"password" validate:"required,min=8"`
	PasswordConfirm string `gorm:"-" form:"password_confirm" validate:"required,eqfield=Password"`
	Admin           bool   `gorm:"default:false" form:"admin" validate:"boolean"`
	SessionKey      string `gorm:""`
}

func (u *User) Create(db *gorm.DB) error {
	// If user with this username is found return err
	if u.IsFound(db) {
		return ErrUserNotUniq
	}
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) IsFound(db *gorm.DB) bool {
	return !errors.Is(u.GetByUsername(db, u.Email), gorm.ErrRecordNotFound)
}

func (u *User) Validate(v *validators.Validator) error {
	if !v.ValidateInputs(u.Email, u.Password, u.PasswordConfirm) {
		return fiber.ErrInternalServerError
	}
	if err := v.SwitchUserValidate(u); err != nil {
		return err
	}
	// Hash password and do u.Password = password
	if u.hashPassword() != nil {
		return ErrPasswordEncrypt
	}
	return nil
}

func (u *User) hashPassword() error {
	password, err := utils.GenHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}

func (u *User) ComparePassword(password string) bool {
	if err := utils.ComparePassword(u.Password, password); err == nil {
		return true
	}
	return false
}

func GetAllUsers(db *gorm.DB) []User {
	users := []User{}
	db.Find(&users)
	return users
}

func (u *User) Get(db *gorm.DB) error {
	return db.First(u, u.ID).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Unscoped().Delete(u).Error
}

func (u *User) GetByUsername(db *gorm.DB, username string) error {
	u.ID = 0
	return db.Where("email = ?", username).First(&u).Error
	// return db.Where("email = ?", username).Take(&u).Error
}
