package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	model "github.com/jacobd39/admsys/models"
)

//GenerateJWT nos proporciona un método para generar un JSON Web Token
func GenerateJWT(t model.User) (string, error) {
	myKey := []byte("kbjnfqfsfy79")

	payload := jwt.MapClaims{
		"dni":   	 t.Dni,
		"email":	 t.Email,
		"idProfile": t.IDProfile,
		"status":    t.Status,
		"name1":     t.Name1,
		"name2":     t.Name2,
		"lastName1": t.LastName1,
		"lastName2": t.LastName2,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
