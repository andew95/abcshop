package abcPassword

import (
	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) string {
// 	pwd := []byte(password)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(hashedPassword)
// }

// func ComparePasswords(hashedPwd string, plainPwd string) bool {
// 	byteHash := []byte(hashedPwd)
// 	bytePlain := []byte(plainPwd)
// 	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
// 	return err == nil
// }

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
