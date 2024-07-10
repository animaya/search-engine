package db

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"-"`
	IsAdmin   bool       `gorm:"default:false" json:"isAdmin"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt time.Time  `json:"updateddAt"`
}

// func (User) TableName() string {
// 	return "public.users"
// }

func (u *User) CreateAdmin() error {
	user := User{
		Email:    "your email",
		Password: "your password",
		IsAdmin:  true,
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		fmt.Println("Error generating password:", err)
		return errors.New("error creating password")
	}

	user.Password = string(password)

	if err := DBConn.Create(&user).Error; err != nil {
		fmt.Println("Error creating user:", err)
		return errors.New("error creating user")
	}
	fmt.Println("Admin user created successfully with email:", u.Email)
	fmt.Println("Stored hashed password:", u.Password) // Log hashed password for verification
	return nil
}

func (u *User) LoginAsAdmin(email string, password string) (*User, error) {
	if err := DBConn.Where("email = ? AND is_admin = ?", email, true).First(&u).Error; err != nil {
		fmt.Println("User not found with email:", email)
		return nil, errors.New("user not found")
	}
	fmt.Println("User found:", u.Email)
	fmt.Println("Stored hashed password for user:", u.Password) // Log stored hashed password

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		fmt.Println("Password comparison failed for user:", u.Email)
		fmt.Println("Provided password:", password) // Log provided password
		return nil, errors.New("invalid password")

	}
	fmt.Println("Password comparison succeeded for user:", u.Email)
	return u, nil
}
