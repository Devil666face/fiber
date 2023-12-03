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

// func (user *User) Set(username string, password string, admin bool) {
// 	user.Username = username
// 	user.Password = password
// 	user.Admin = admin
// }

// func CreateUser(user *User) *gorm.DB {
// 	return database.DB.Create(user)
// }

// func UpdateUser(user *User) *gorm.DB {
// 	return database.DB.Save(user)
// }

// func DeleteUser(user *User) *gorm.DB {
// 	return database.DB.Unscoped().Delete(user)
// }

// func GetUserByUsername(dest *User, username string) *gorm.DB {
// 	return database.DB.Where("username= ?", username).Take(&dest)
// }

// func GetAllUsers() []User {
// 	var users []User
// 	database.DB.Find(&users)
// 	return users
// }

// type UserForm struct {
// 	Username        string `form:"username"`
// 	Password        string `form:"password"`
// 	PasswordConfirm string `form:"password_confirm"`
// 	Admin           string `form:"admin"`
// }

// func (form *UserForm) IsAdmin() bool {
// 	if form.Admin != "" {
// 		return true
// 	}
// 	return false
// }

// func (form *UserForm) IsEmptyUsername() (string, bool) {
// 	if form.Username == "" {
// 		return "Username is required.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsMatch() (string, bool) {
// 	if form.Password != form.PasswordConfirm {
// 		return "The passwords don't match.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsEmpty() (string, bool) {
// 	if form.Password == "" || form.PasswordConfirm == "" {
// 		return "Password is required.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsShort() (string, bool) {
// 	if len([]rune(form.Password)) < utils.PASSWORDLEN {
// 		return fmt.Sprintf("The minimum len of password is %d", utils.PASSWORDLEN), true
// 	}
// 	return "", false
// }

// func (form *UserForm) CheckPasswordForCreate() (string, bool) {
// 	if message, ok := form.IsPasswordsMatch(); ok {
// 		return message, ok
// 	}
// 	if message, ok := form.IsPasswordsEmpty(); ok {
// 		return message, ok
// 	}
// 	if message, ok := form.IsPasswordsShort(); ok {
// 		return message, ok
// 	}
// 	return "", false
// }
