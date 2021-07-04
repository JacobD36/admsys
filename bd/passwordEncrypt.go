package bd

import "golang.org/x/crypto/bcrypt"

//PasswordEncrypt encripta el password que se le pasa como parámetro
func PasswordEncrypt(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}
