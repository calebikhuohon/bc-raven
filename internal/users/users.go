package users

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"-"`
}

func Create(ctx context.Context, client firestore.Client, data *User) error  {
	hashedPassword, err := HashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = hashedPassword
	_, err = client.Collection("users").Doc(data.UserName).Set(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string, ctx context.Context, client firestore.Client) (string, error) {
	dsnap, err := client.Collection("users").Doc(username).Get(ctx)
	if err != nil {
		return "", err
	}

	var d User
	if err := dsnap.DataTo(&d); err != nil {
		return "", err
	}
	return d.ID, nil
}

func (user *User) Authenticate(ctx context.Context, client firestore.Client) bool  {
	dsnap, err := client.Collection("users").Doc(user.UserName).Get(ctx)
	if err != nil {
		return false
	}

	var d User
	if err := dsnap.DataTo(&d); err != nil {
		return false
	}

	return CheckPasswordHash(user.Password, d.Password)
}
