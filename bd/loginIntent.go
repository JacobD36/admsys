package bd

import (
	model "github.com/jacobd39/admsys/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginIntent realiza el intento de Login
func LoginIntent(dni string, password string) (model.User, bool, int) {
	user, found, _ := UserVerification(dni)

	if found == false {
		return user, false, 1
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false, 2
	}

	return user, true, 0
}
